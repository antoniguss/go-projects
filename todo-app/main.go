package main

import (
	"fmt"
)

type todo struct {
	ID          int
	Description string
	IsCompleted bool
}

type TodoManager struct {
	nextID int
	todos  []todo
}

func (tm *TodoManager) AddTodo(desc string) (id int) {

	newTodo := todo{
		ID:          tm.nextID,
		Description: desc,
		IsCompleted: false,
	}
	tm.todos = append(tm.todos, newTodo)

	tm.nextID++

	return newTodo.ID

}

func (tm *TodoManager) RemoveTodo(id int) error {

	for i, t := range tm.todos {
		if t.ID == id {
			tm.todos = append(tm.todos[:i], tm.todos[i+1:]...)
			return nil
		}

	}

	return fmt.Errorf("no todo with id %d", id)
}

func (tm *TodoManager) MarkCompleted(id int) error {
	for i, t := range tm.todos {
		if t.ID == id {
			tm.todos[i].IsCompleted = true
			return nil
		}
	}

	return fmt.Errorf("no todo with id %d", id)

}

func NewTodoManager() *TodoManager {
	return &TodoManager{
		nextID: 0,
		todos:  make([]todo, 0),
	}
}

func main() {
	tm := NewTodoManager()

	tm.AddTodo("Buy dog snacks")
	id1 := tm.AddTodo("Do the laundry")
	id2 := tm.AddTodo("Push to Github")

	fmt.Println(tm)

	err := tm.RemoveTodo(id1)
	if err != nil {
		fmt.Println(err)
	}

	err = tm.MarkCompleted(id2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(tm)
}
