package main

import (
	"fmt"
)

func main() {
	tm := NewTodoManager()

	tm.AddTodo("Buy dog snacks")
	id1 := tm.AddTodo("Do the laundry, only white")
	id2 := tm.AddTodo("Push to Github")

	tm.MarkCompleted(id1)
	tm.SaveToFile("./todos.csv")
	err := tm.RemoveTodo(id1)
	if err != nil {
		fmt.Println(err)
	}

	err = tm.MarkCompleted(id2)
	if err != nil {
		fmt.Println(err)
	}

	tm.DisplayTodos(true)
	fmt.Println()

	err = tm.ReadFromFile("./todos.csv", true)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	tm.DisplayTodos(true)
	fmt.Println()

	err = tm.ReadFromFile("./todos_wrong.csv", true)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	tm.DisplayTodos(true)

}
