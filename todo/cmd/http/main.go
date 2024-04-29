package main

import (
	"log"
	"net/http"
	"todo/pkg/handlers"
	"todo/pkg/service"
)

func main() {

	handlers.NewValidator()

	todoService := service.TodoService{}
	todoHandler := handlers.NewTodoHandler(&todoService)

	http.HandleFunc("GET /todos", todoHandler.GetAllTodosView)

	http.HandleFunc("GET /api/todos", todoHandler.GetAllTodos)
	http.HandleFunc("POST /api/todos", todoHandler.CreateTodo)
	http.HandleFunc("GET /api/todos/{id}", todoHandler.GetTodoById)
	http.HandleFunc("PUT /api/todos/{id}", todoHandler.UpdateTodoById)
	http.HandleFunc("DELETE /api/todos/{id}", todoHandler.DeleteTodoById)

	log.Println("server started at :5000")
	log.Fatalf("server error: %v", http.ListenAndServe(":5000", nil))
}
