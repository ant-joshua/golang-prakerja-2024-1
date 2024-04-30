package models

type CreateProjectRequest struct {
	Title       string              `json:"title" validate:"required"`
	Description string              `json:"description"`
	Todos       []CreateTodoRequest `json:"todos" validate:"required,dive"`
}
