package handlers

import (
	"APIGateway/models"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

// GetNewsList возвращает список новостей, поддерживая фильтрацию по заголовку
func GetNewsList(w http.ResponseWriter, r *http.Request) {
	// Добавляем пагинацию
	titleQuery := r.URL.Query().Get("title")
	pageStr := r.URL.Query().Get("page")
	sizeStr := r.URL.Query().Get("size")

	// Конвертируем параметры пагинации в числа, используем значения по умолчанию при ошибке
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	size, err := strconv.Atoi(sizeStr)
	if err != nil || size < 1 {
		size = 10 // указываем размер страницы по умолчанию
	}

	news := filterNewsByTitle(titleQuery)
	start := (page - 1) * size
	end := start + size

	if start > len(news) {
		start = len(news)
	}

	if end > len(news) {
		end = len(news)
	}

	paginatedNews := news[start:end]

	// Добавляем информацию о пагинации в ответ
	response := struct {
		News  []models.NewsShortDetailed `json:"news"`
		Page  int                        `json:"page"`
		Size  int                        `json:"size"`
		Total int                        `json:"total"`
	}{
		News:  paginatedNews,
		Page:  page,
		Size:  size,
		Total: len(news),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

// filterNewsByTitle фильтрует новости по заголовкам
func filterNewsByTitle(query string) []models.NewsShortDetailed {
	allNews := []models.NewsShortDetailed{
		{ID: 1, Title: "Первая новость"},
		{ID: 2, Title: "Вторая новость"},
	}
	filteredNews := []models.NewsShortDetailed{}
	for _, news := range allNews {
		if contains(news.Title, query) {
			filteredNews = append(filteredNews, news)
		}
	}
	return filteredNews
}

// contains проверяет, содерджится ли подстрока в строке
func contains(str, substr string) bool {
	return strings.Contains(str, substr)
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
