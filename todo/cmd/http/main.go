package main

import (
	"log"
	"net/http"
	service "todo/pkg"
	"todo/pkg/handlers"
)

func main() {

	todoService := service.TodoService{}
	todoHandler := handlers.NewTodoHandler(&todoService)

	http.HandleFunc("GET /todos", todoHandler.GetAllTodos)

	log.Println("server started at :5000")
	log.Fatalf("server error: %v", http.ListenAndServe(":5000", nil))
}
