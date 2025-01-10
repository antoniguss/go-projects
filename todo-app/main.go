/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/antoniguss/go-projects/todo-app/cmd"
)

func main() {

	cmd.Execute()

}

// Specify a command to execute:
//
//	add <description>       Add a new TODO with the given description
//	display [-a, --all]         Display TODOs (--all shows completed items)
//	remove <id>             Remove TODO with specified ID
//	complete <id> [-u, --undo]  Mark/unmark TODO as completed (--undo to mark incomplete)
// func executeCommand(command string, args []string) error {
// 	printfDebug("Executing: %s, args: %v\n", command, args)

// 	switch command {
// 	case "add":
// 		if len(args) != 1 || len(args[0]) == 0 {
// 			return fmt.Errorf("'add' command requires a non-empty description")
// 		}
// 		id := todoManager.AddTodo(args[0])
// 		printfDebug("Added TODO with id %d\n", id)
// 		return nil

// 	case "display":
// 		if len(args) > 1 {
// 			return fmt.Errorf("'display' command accepts only one optional flag: --all or -a")
// 		}
// 		if len(args) == 0 {
// 			todoManager.DisplayTodos(false)
// 			return nil
// 		}
// 		if args[0] == "--all" || args[0] == "-a" {
// 			todoManager.DisplayTodos(true)
// 			return nil
// 		}
// 		return fmt.Errorf("invalid flag for 'display' command. Use --all or -a")

// 	case "remove":
// 		if len(args) != 1 {
// 			return fmt.Errorf("'remove' command requires a todo ID")
// 		}
// 		id, err := strconv.Atoi(args[0])
// 		if err != nil {
// 			return fmt.Errorf("invalid todo ID: %s (must be a number)", args[0])
// 		}
// 		if err := todoManager.RemoveTodo(id); err != nil {
// 			return fmt.Errorf("failed to remove todo: %v", err)
// 		}
// 		return nil

// 	case "complete":
// 		if len(args) == 0 {
// 			return fmt.Errorf("'complete' command requires a todo ID")
// 		}
// 		if len(args) > 2 {
// 			return fmt.Errorf("'complete' command accepts only a todo ID and an optional flag: --undo or -u")
// 		}

// 		id, err := strconv.Atoi(args[0])
// 		if err != nil {
// 			return fmt.Errorf("invalid todo ID: %s (must be a number)", args[0])
// 		}

// 		undo := false
// 		if len(args) == 2 {
// 			if args[1] == "--undo" || args[1] == "-u" {
// 				undo = true
// 			} else {
// 				return fmt.Errorf("invalid flag for 'complete' command. Use --undo or -u")
// 			}
// 		}

// 		if err := todoManager.SetCompleted(id, !undo); err != nil {
// 			return fmt.Errorf("failed to update todo completion status: %v", err)
// 		}
// 		return nil

// 	default:
// 		return fmt.Errorf("unknown command '%s'. Run without arguments to see usage", command)
// 	}
// }

// func printDebug(msg string) {
// 	if *debug {
// 		fmt.Printf("[DEBUG]: %s\n", msg)
// 	}

// }

// func printfDebug(format string, a ...any) {
// 	if *debug {
// 		format = "[DEBUG]: " + format

// 		fmt.Printf(format, a...)
// 	}
// }

// func example() {
// 	tm := NewTodoManager()

// 	tm.AddTodo("Buy dog snacks")
// 	id1 := tm.AddTodo("Do the laundry, only white")
// 	id2 := tm.AddTodo("Push to Github")

// 	tm.SetCompleted(id1, true)
// 	tm.Export("./todos.csv")
// 	err := tm.RemoveTodo(id1)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	err = tm.SetCompleted(id2, true)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	tm.DisplayTodos(true)
// 	fmt.Println()

// 	err = tm.Import("./todos.csv", true)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}
// 	tm.DisplayTodos(true)
// 	fmt.Println()

// 	err = tm.Import("./todos_wrong.csv", true)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}
// 	tm.DisplayTodos(true)
// }
