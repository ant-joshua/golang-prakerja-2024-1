package models

import "time"

type CreateTodoRequest struct {
	Title     string    `json:"title" validate:"required"`
	ExpiredAt time.Time `json:"expired_at" validate:"required"`
}

type UpdateTodoRequest struct {
	Title     *string    `json:"title"`
	Done      *bool      `json:"done"`
	ExpiredAt *time.Time `json:"expired_at"`
}

type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	ExpiredAt time.Time `json:"expired_at"`
}

type Todos []Todo

const (
	ErrNoRecord  = "no record found"
	ErrInvalidID = "invalid id"
)
