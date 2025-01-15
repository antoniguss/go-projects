package storage

import "github.com/antoniguss/go-projects/todo-app/types"

type Storage interface {
	Setup() error
	Close() error
	AddTodo(description string) error
	GetAllTodos() ([]types.Todo, error)
	SetCompleted(id int, completed bool) error
	RemoveTodo(id int) error
}
