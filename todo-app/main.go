package main

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

type todo struct {
	ID            int
	Description   string
	TimeAdded     time.Time
	TimeComplated time.Time
	IsCompleted   bool
}

type TodoManager struct {
	nextID int
	todos  []todo
}

func (tm *TodoManager) AddTodo(desc string) (id int) {

	newTodo := todo{
		ID:          tm.nextID,
		Description: desc,
		TimeAdded:   time.Now(),
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
			tm.todos[i].TimeComplated = time.Now()
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

func (tm *TodoManager) DisplayTodos(all bool) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)

	if all {
		fmt.Fprintf(w, "ID\tDescription\tCompleted\tAdded\t\n")

		for _, t := range tm.todos {

			fmt.Fprintf(w, "%d\t%s\t%v\t%v ago\t\n",
				t.ID,
				t.Description,
				t.IsCompleted,
				time.Now().Sub(t.TimeAdded),
			)

		}
	} else {

		fmt.Fprintf(w, "ID\tDescription\n")

		for _, t := range tm.todos {
			fmt.Fprintf(w, "%d\t%s\n",
				t.ID, t.Description)

		}
	}

	w.Flush()

}

func main() {
	tm := NewTodoManager()

	tm.AddTodo("Buy dog snacks")
	id1 := tm.AddTodo("Do the laundry")
	id2 := tm.AddTodo("Push to Github")

	err := tm.RemoveTodo(id1)
	if err != nil {
		fmt.Println(err)
	}

	err = tm.MarkCompleted(id2)
	if err != nil {
		fmt.Println(err)
	}
	tm.DisplayTodos(true)

}
