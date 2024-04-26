package main

import (
	"fmt"
	"time"
	"todo/pkg/models"
	"todo/pkg/service"
)

func todoExample() {
	todoService := service.TodoService{}

	fmt.Println("all todos data", todoService.GetAllTodos())

	newTodo := todoService.AddNewTodo(models.Todo{
		ID:        3,
		Title:     "Third Todo",
		Done:      false,
		ExpiredAt: time.Now(),
	})

	fmt.Println("new todo added", newTodo)

	fmt.Println("see all todos", todoService.GetAllTodos())

	fmt.Println("search todo with id 2")
	todo, err := todoService.GetTodoById(2)
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("todo", todo)
	}

	fmt.Println("search todo with id 4")
	todo, err = todoService.GetTodoById(4)
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("todo", todo)
	}

	fmt.Println("update todo with id 2")
	updatedTodo, err := todoService.UpdateTodoById(2, models.Todo{
		ID:        2,
		Title:     "Second Todo Updated",
		Done:      true,
		ExpiredAt: time.Now(),
	})

	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("todo updated", updatedTodo)
	}

	fmt.Println("see all todos", todoService.GetAllTodos())

	fmt.Println("delete todo with id 3")

	err = todoService.DeleteTodoById(3)
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("todo deleted")
	}

	fmt.Println("see all todos", todoService.GetAllTodos())
}

func main() {
	todoExample()
}
