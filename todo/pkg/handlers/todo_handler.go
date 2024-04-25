package handlers

import (
	"encoding/json"
	"net/http"
	service "todo/pkg"
)

type TodoHandler struct {
	service *service.TodoService
}

func NewTodoHandler(service *service.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos := h.service.GetAllTodos()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(todos)
}
