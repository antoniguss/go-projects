package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	usage = `Specify a command to execute:
    - add <desc>: Add a TODO with a description
    - display: Display all TODOs
    - remove <id>: Remove TODO with given id
    - complete <id>: Mark TODO with given id as completed
    - import <filepath>: Import TODOs from a given .csv file
    - export <filepath>: Export TODOs to a given .csv file`

	debug       = flag.Bool("debug", false, "log out all the debug information")
	todoManager = TodoManager{}
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

	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)

	}

	command := args[0]

	err := executeCommand(command, args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

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
		return nil
	case "complete":
		return nil
	case "import":
		return nil
	case "export":
		return nil
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
