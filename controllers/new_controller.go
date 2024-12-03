package controllers

import (
	"APIGateway/models"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

type results struct {
	data interface{}
	err  error
}

func GetNewsDetails(w http.ResponseWriter, r *http.Request) {
	// Создвем канал для результатов
	result := make(chan results, 2)
	var wg sync.WaitGroup

	// Извлекаем ID новости из параметров пути
	vars := mux.Vars(r)
	newsIDStr, ok := vars["id"]
	if !ok {
		http.Error(w, "News ID is required", http.StatusBadRequest)
		return
	}

	NewsID, err := strconv.Atoi(newsIDStr)
	if err != nil {
		http.Error(w, "Invalid News ID", http.StatusBadRequest)
		return
	}

	// Увеличиваем счетчик WaitGroup
	wg.Add(2)

	// Запрос к агрегатору новостей
	go func() {
		defer wg.Done()
		newsDetalis, err := fetchNewsDetails(NewsID)
		result <- results{data: newsDetalis, err: err}
	}()

	// Запрос к сервису комментариев
	go func() {
		defer wg.Done()
		comments, err := fetchComments(NewsID)
		result <- results{data: comments, err: err}
	}()

	// Ожидаем заврешения всех запросов
	wg.Wait()
	close(result)

	//Обработка результатов
	var news models.NewsFullDetailed
	var comments []models.Comment
	for res := range result {
		if res.err != nil {
			http.Error(w, "Ошибка при получении данных", http.StatusInternalServerError)
			return
		}
		switch v := res.data.(type) {
		case models.NewsFullDetailed:
			news = v
		case []models.Comment:
			comments = v
		}

	}

	// Формируем и отправляем ответ
	response := struct {
		News     models.NewsFullDetailed `json:"news"`
		Comments []models.Comment        `json:"comments"`
	}{news, comments}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

// fetchNewsDetails получает детальную информацию о новости по ее ID.
func fetchNewsDetails(newsID int) (models.NewsFullDetailed, error) {
	// Логика получения новостей из бд
	return models.NewsFullDetailed{
		ID:      newsID,
		Title:   "Пример заголовка",
		Content: "Полное содержание новости.",
	}, nil
}

// fetchComments получает комментарии к новости по ее ID.
func fetchComments(newsID int) ([]models.Comment, error) {
	// Логика получения комментариев из бд

	comments := []models.Comment{
		{ID: 1, NewsID: newsID, Text: "Первый комментарий", IsModerated: true},
		{ID: 2, NewsID: newsID, Text: "Второй комментарий", IsModerated: true},
	}
	return comments, nil
}
