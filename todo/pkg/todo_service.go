package service

import (
	"errors"
	"slices"
	"todo/pkg/models"
)

var todoList = models.Todos{
	models.Todo{
		ID:    1,
		Title: "First Todo",
		Done:  false,
	},
	models.Todo{
		ID:    2,
		Title: "Second Todo",
		Done:  true,
	},
}

func GetAllTodos() models.Todos {

	return todoList
}

func AddNewTodo(newTodo models.Todo) models.Todo {

	todoList = append(todoList, newTodo)
	return newTodo
}

func GetTodoById(id int) (*models.Todo, error) {
	var todo *models.Todo

	for _, t := range todoList {
		if t.ID == id {
			todo = &t
			break
		}
	}
	if todo == nil {
		return nil, errors.New(models.ErrNoRecord)
	}

	return todo, nil
}

func UpdateTodoById(id int, updatedTodo models.Todo) (*models.Todo, error) {
	var todo *models.Todo
	for i, t := range todoList {
		if t.ID == id {
			todoList[i] = updatedTodo
			todo = &updatedTodo
			break
		}
	}
	if todo == nil {
		return nil, errors.New(models.ErrNoRecord)
	}

	return todo, nil
}

func DeleteTodoById(id int) error {
	for i, t := range todoList {
		if t.ID == id {
			deleted := slices.Delete(todoList, i, i+1)

			if deleted == nil {
				return errors.New(models.ErrNoRecord)

			}

			todoList = deleted
			// todoList = append(todoList[:i], todoList[i+1:]...)
			return nil
		}
	}
	return errors.New(models.ErrNoRecord)
}
