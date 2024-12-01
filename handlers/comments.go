package handlers

import (
	"APIGateway/models"
	"encoding/json"
	"net/http"
)

// AddComment добавляет новый комментарий
func AddComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, "Ошибка при деодировании комментария", http.StatusBadRequest)
		return
	}

	//Сохраняем комментарий в БД (зашлушка)
	comment.ID = 1
	comment.IsModerated = true

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comment)
}

// GetCommentsByNwesID возвращает все комментарии по ID нвовости
func GetCommentsByNewsID(w http.ResponseWriter, r *http.Request) {
	comments := []models.Comment{
		{ID: 1, NewsID: 1, Text: "Первый комментарий", IsModerated: true},
		{ID: 2, NewsID: 1, Text: "Второй комментарий", IsModerated: true},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comments)
}
