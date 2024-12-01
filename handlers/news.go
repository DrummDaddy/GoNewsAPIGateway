package handlers

import (
	"APIGateway/models"
	"encoding/json"
	"net/http"
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
