package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/pkg/service"
)

var todoService = service.TodoService{}
var todoHandler = NewTodoHandler(&todoService)

func TestGetAllTodoView(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	w := httptest.NewRecorder()
	todoHandler.GetAllTodosView(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status OK, got %v", w.Code)
	}
}

func TestExample(t *testing.T) {
	fmt.Println("Test Example")
}
