package main

import (
	"github.com/antoniguss/go-projects/todo-app/cmd"
	_ "github.com/antoniguss/go-projects/todo-app/cmd"
	_ "github.com/antoniguss/go-projects/todo-app/storage"
)

func main() {
	cmd.Execute()
}
