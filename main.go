package main

import (
	"APIGateway/database"
	"APIGateway/handlers"
	"net/http"

	"log"

	//"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	//Инициализация соединения с БД
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}
	defer db.Close()
	//Создаем новый маршрутизатор
	r := mux.NewRouter()

	// Определяем маршруты
	r.HandleFunc("/news", handlers.GetNewsList).Methods("GET")
	r.HandleFunc("/news/filter", handlers.FiltersNews).Methods("GET")
	r.HandleFunc("/news/{id}", handlers.GetNewsDetails).Methods("GET")
	r.HandleFunc("/comments", handlers.AddComment).Methods("POST")
	r.HandleFunc("/news/{id}/comments", handlers.GetCommentsByNewsID).Methods("GET")

	//Запускаем сервер
	log.Println("Запуск сервера на http://localhost:8080/")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/news", NewsHandler)

	// Оборачивание обработчиков в middleware
	chain := RequestIDMiddleware(LoggingMiddleware(mux))

	http.ListenAndServe(":8080", chain)

}
