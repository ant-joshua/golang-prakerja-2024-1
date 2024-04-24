package main

import (
	"fmt"
	"time"
	service "todo/pkg"
	"todo/pkg/models"
)

func main() {
	fmt.Println("all todos data", service.GetAllTodos())

	newTodo := service.AddNewTodo(models.Todo{
		ID:        3,
		Title:     "Third Todo",
		Done:      false,
		ExpiredAt: time.Now(),
	})

	fmt.Println("new todo added", newTodo)

	fmt.Println("see all todos", service.GetAllTodos())

	fmt.Println("search todo with id 2")
	todo, err := service.GetTodoById(2)
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("todo", todo)
	}

	fmt.Println("search todo with id 4")
	todo, err = service.GetTodoById(4)
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("todo", todo)
	}

	fmt.Println("update todo with id 2")
	updatedTodo, err := service.UpdateTodoById(2, models.Todo{
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

	fmt.Println("see all todos", service.GetAllTodos())

	fmt.Println("delete todo with id 3")

	err = service.DeleteTodoById(3)
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("todo deleted")
	}

	fmt.Println("see all todos", service.GetAllTodos())

}
