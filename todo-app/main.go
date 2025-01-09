package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	usage = `Specify a command to execute:
    - add <desc>: Add a TODO with a description
    - display: Display all TODOs
    - remove <id>: Remove TODO with given id
    - complete <id>: Mark TODO with given id as completed`
	// - import <filepath>: Import TODOs from a given .csv file
	// - export <filepath>: Export TODOs to a given .csv file`

	todosPath = "./todos.csv"

	debug       = flag.Bool("debug", false, "log out all the debug information")
	todoManager TodoManager
)

func main() {
	// example()
	//--- CLI FUNCTIONALITY ---
	// addTodo(description string)
	// displayTodos()
	// removeTodo(id int)
	// complete(id int)
	// import(filePath string)
	// export(filePath string)

	flag.Parse()
	args := flag.Args()

	todoManager = TodoManager{}
	err := todoManager.Import(todosPath, false)
	if err != nil {
		printDebug(err.Error())
	}

	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)

	}

	command := args[0]

	err = executeCommand(command, args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	todoManager.Export("./todos.csv")

}

// `Specify a command to execute:
//   - add <desc>: Add a TODO with a description
//   - display: Display all TODOs
//   - remove <id>: Remove TODO with given id
//   - complete <id>: Mark TODO with given id as completed
//   - import <filepath>: Import TODOs from a given .csv file
//   - export <filepath>: Export TODOs to a given .csv file`
func executeCommand(command string, args []string) error {

	printfDebug("Executing: %s, args: %v\n", command, args)

	switch command {
	case "add":
		if len(args) != 1 {
			return fmt.Errorf("add command expects 1 non-empty string argument")
		}

		desc := args[0]
		if len(desc) == 0 {
			return fmt.Errorf("add command expects 1 non-empty string argument")
		}
		id := todoManager.AddTodo(desc)
		printfDebug("Added TODO with id %d\n", id)

		return nil
	case "display":
		if len(args) > 0 {
			return fmt.Errorf("display command doesn't expect any arguments")
		}

		todoManager.DisplayTodos(true)
		return nil
	case "remove":

		if len(args) != 1 {
			return fmt.Errorf("remove command expects 1 non-empty string argument")
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("remove command expects 1 non-empty string argument, (%s)", err.Error())

		}

		err = todoManager.RemoveTodo(id)
		if err != nil {
			return err
		}
		return nil
	case "complete":
		if len(args) != 1 {
			return fmt.Errorf("complete command expects 1 non-empty string argument")
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("complete command expects 1 non-empty string argument, (%s)", err.Error())

		}

		err = todoManager.MarkCompleted(id)
		if err != nil {
			return err
		}
		return nil
	// case "import":
	// 	return nil
	// case "export":
	// 	return nil
	default:
		return fmt.Errorf("invalid command '%s'", command)
	}

}

func printDebug(msg string) {
	if *debug {
		fmt.Printf("[DEBUG]: %s\n", msg)
	}

}

func printfDebug(format string, a ...any) {
	if *debug {
		format = "[DEBUG]: " + format

		fmt.Printf(format, a...)
	}
}

func example() {
	tm := NewTodoManager()

	tm.AddTodo("Buy dog snacks")
	id1 := tm.AddTodo("Do the laundry, only white")
	id2 := tm.AddTodo("Push to Github")

	tm.MarkCompleted(id1)
	tm.Export("./todos.csv")
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

	err = tm.Import("./todos.csv", true)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	tm.DisplayTodos(true)
	fmt.Println()

	err = tm.Import("./todos_wrong.csv", true)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	tm.DisplayTodos(true)
}
