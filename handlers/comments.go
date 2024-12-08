package handlers

import (
	"APIGateway/models"
	"bytes"
	"encoding/json"
	"fmt"
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

	//Вызов сервиса цензурирования
	if err := checkCommentWithCensorshipService(comment.Content); err != nil {
		http.Error(w, "Comment failed validation", http.StatusBadRequest)
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

func checkCommentWithCensorshipService(content string) error {
	reqBody, err := json.Marshal(map[string]string{"comment": content})
	if err != nil {
		return err
	}

	resp, err := http.Post("http://localhost:8080/validate", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("comment was not validated")
	}

	return nil
}
