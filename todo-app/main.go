package main

import (
	"github.com/antoniguss/go-projects/todo-app/cmd"
	_ "github.com/antoniguss/go-projects/todo-app/cmd"
	_ "github.com/antoniguss/go-projects/todo-app/manager"
	_ "github.com/antoniguss/go-projects/todo-app/storage"
)

func main() {

	// err := storage.Setup()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// err = storage.AddTodo(manager.Todo{
	// 	Description: "Do Laundry",
	// })
	//
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// todos, err := storage.GetAllTodos()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("todos: %v\n", todos)
	cmd.Execute()
	// err = storage.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
