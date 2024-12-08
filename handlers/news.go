package handlers

import (
	"APIGateway/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
)

// GetNewsList возвращает список новостей
func GetNewsList(w http.ResponseWriter, r *http.Request) {
	news := []models.NewsShortDetailed{
		{ID: 1, Title: "Первая новость"},
		{ID: 2, Title: "Вторая новость"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news)
}

// FiltersNews фильтрует новости
func FiltersNews(w http.ResponseWriter, r *http.Request) {
	filteredNews := []models.NewsShortDetailed{
		{ID: 1, Title: "Отфильтрованная новость"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filteredNews)
}

// GetNewsDetails возвращает детальную информацию о новости
func GetNewsDetails(w http.ResponseWriter, r *http.Request) {
	news := models.NewsFullDetailed{
		ID:      1,
		Title:   "Детали новости",
		Content: "Это полное содержание новости.",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news)
}

type CombinedResponse struct {
	News     models.News      `json:"news"`
	Comments []models.Comment `json:"comments"`
}

func HandleNewsRequest(w http.ResponseWriter, r *http.Request) {
	var (
		news     models.News
		comments []models.Comment
		wg       sync.WaitGroup
	)

	// Асинхронные вызовы
	wg.Add(1)
	go func() {
		defer wg.Done()
		comments = fetchComments()
	}()

	wg.Wait()

	response := CombinedResponse{
		News:     news,
		Comments: comments,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func fetchNews() models.News {
	return models.News{ID: 1, Title: "Sample News"}
}

func fetchComments() []models.Comment {
	return []models.Comment{{NewsID: 1, Text: "Great article!"}}
}

// Обработчик запросов
func NewsHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	searchQuery := queryParams.Get("s")
	page := queryParams.Get("page")

	// Подготавливаем запрос для агрегатора, добавляя новые параметры
	aggregatorURL := fmt.Sprintf("http://aggregator-service/news?s=%s&page=%s",
		url.QueryEscape(searchQuery), url.QueryEscape(page))

	// Выполняем запрос к агрегатору
	resp, err := http.Get(aggregatorURL)
	if err != nil {
		http.Error(w, "Error reaching the news aggregator", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Перенаправление ответа от агрегатора клиенту
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Error from aggregator", http.StatusBadGateway)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}
