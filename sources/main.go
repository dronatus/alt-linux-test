package main

import (
	"fmt"
	"log"
	"net/http"
	"simple-quiz/handlers"
	"time"
)

func main() {
	// Инициализация обработчиков
	if err := handlers.Init(); err != nil {
		log.Fatalf("Failed to initialize handlers: %v", err)
	}

	// Настройка маршрутов
	http.HandleFunc("/", handlers.QuizHandler)
	http.HandleFunc("/result", handlers.ResultHandler)

	// Обслуживание статических файлов
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Настройка сервера
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	fmt.Println("Сервис 'Какой ты Смешарик' запускается...")
	fmt.Println("Доступен по адресу: http://localhost:8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
