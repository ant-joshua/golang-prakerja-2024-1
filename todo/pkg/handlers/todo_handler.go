package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
	"todo/pkg/models"
	"todo/pkg/service"
	"todo/web"

	"github.com/go-playground/validator/v10"
)

type TodoHandler struct {
	service *service.TodoService
}

func NewTodoHandler(service *service.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) GetAllTodosView(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	var filepath = path.Join(web.Root, "todos.html")
	var tmpl, err = template.ParseFiles(filepath)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"title": "Learning Golang Web",
		"todos": h.service.GetAllTodos(),
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (h *TodoHandler) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos := h.service.GetAllTodos()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(todos)
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// var createTodoRequest models.CreateTodoRequest
	var createTodoRequest models.CreateProjectRequest
	json.NewDecoder(r.Body).Decode(&createTodoRequest)

	err := Validate.Struct(createTodoRequest)

	if err != nil {

		customErr := make(map[string]string)

		// errs := err.(validator.ValidationErrors)

		var errs validator.ValidationErrors

		if errors.As(err, &errs) {
			w.WriteHeader(http.StatusBadRequest)

			validationErr := ErrorValidationHandlerMap(errs)
			validationResp := ErrorValidateResponse(validationErr)

			json.NewEncoder(w).Encode(validationResp)
		}

		// for _, e := range errs {
		// 	fmt.Printf("Error: %v\n", e)
		// 	// fmt.Printf("Error %v\n", e.Type().Key())
		// 	fmt.Printf("Error %v\n", e.Field())

		// }

		fmt.Printf("Custom Error: %v\n", customErr)

		return
	}

	todo := models.Todo{
		ID:    len(h.service.GetAllTodos()) + 1,
		Title: createTodoRequest.Title,
		Done:  false,
	}

	h.service.AddNewTodo(todo)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) GetTodoById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	parseID, err := strconv.Atoi(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid ID"))
		return
	}

	todo, err := h.service.GetTodoById(parseID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) UpdateTodoById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	parseID, err := strconv.Atoi(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid ID"))
		return
	}

	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)

	updateTodo, err := h.service.UpdateTodoById(parseID, todo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updateTodo)
}

func (h *TodoHandler) DeleteTodoById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	parseID, err := strconv.Atoi(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid ID"))
		return
	}

	err = h.service.DeleteTodoById(parseID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Todo Deleted"))
}
