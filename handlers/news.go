package handlers

import (
	"APIGateway/models"
	"encoding/json"
	"net/http"
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
