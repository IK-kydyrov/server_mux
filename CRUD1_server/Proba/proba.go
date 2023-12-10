package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)     // Устанавливаем обработчик для корневого пути
	http.ListenAndServe(":8080", nil) // Запускаем веб-сервер на порту 8080
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, friend!") // Отправляем "Привет, мир!" в ответ на запрос
	fmt.Println("Hello")
}
