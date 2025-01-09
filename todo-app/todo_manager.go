package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

type Todo struct {
	ID            int
	Description   string
	TimeAdded     time.Time
	TimeCompleted time.Time
	IsCompleted   bool
}

type TodoManager struct {
	nextID int
	todos  []Todo
}

func (tm *TodoManager) AddTodo(desc string) int {
	newTodo := Todo{
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
func (tm *TodoManager) SetCompleted(id int, completed bool) error {

	for i, t := range tm.todos {
		if t.ID == id {
			tm.todos[i].IsCompleted = completed
			if completed {
				tm.todos[i].TimeCompleted = time.Now()
			} else {
				tm.todos[i].TimeCompleted = time.Time{} // Zero value
			}
			return nil

		}
	}
	return fmt.Errorf("no todo with id %d", id)
}

func (tm *TodoManager) Import(filePath string, overwrite bool) error {
	if !strings.HasSuffix(filePath, ".csv") {
		return errors.New("path must be to a .csv file")
	}

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Skip header row
	_, err = reader.Read()
	if err != nil {
		return fmt.Errorf("error reading header: %v", err)
	}

	var newTodos []Todo
	newTodos = make([]Todo, 0)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error reading record: %v", err)
		}

		todo, err := parseTodoRecord(record)
		if err != nil {
			return fmt.Errorf("error parsing record: %v", err)
		}

		newTodos = append(newTodos, todo)
		if todo.ID >= tm.nextID {
			tm.nextID = todo.ID + 1
		}
	}

	if overwrite {
		tm.todos = make([]Todo, 0)
		tm.nextID = 0
	}
	tm.todos = append(tm.todos, newTodos...)

	return nil
}

func parseTodoRecord(record []string) (Todo, error) {
	if len(record) != 5 {
		return Todo{}, fmt.Errorf("invalid record length: expected 5, got %d", len(record))
	}

	id, err := strconv.Atoi(record[0])
	if err != nil {
		return Todo{}, fmt.Errorf("invalid ID: %v", err)
	}

	timeAdded, err := time.Parse(time.RFC3339, record[2])
	if err != nil {
		return Todo{}, fmt.Errorf("invalid TimeAdded: %v", err)
	}

	timeCompleted, err := time.Parse(time.RFC3339, record[3])
	if err != nil {
		return Todo{}, fmt.Errorf("invalid TimeCompleted: %v", err)
	}

	isCompleted, err := strconv.ParseBool(record[4])
	if err != nil {
		return Todo{}, fmt.Errorf("invalid IsCompleted: %v", err)
	}

	return Todo{
		ID:            id,
		Description:   record[1],
		TimeAdded:     timeAdded,
		TimeCompleted: timeCompleted,
		IsCompleted:   isCompleted,
	}, nil
}

func (tm *TodoManager) Export(filePath string) error {
	if len(tm.todos) == 0 {
		return nil
	}

	if !strings.HasSuffix(filePath, ".csv") {
		return errors.New("path must be to a .csv file")
	}

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write headers
	headers := []string{"ID", "Description", "TimeAdded", "TimeCompleted", "IsCompleted"}
	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("error writing headers: %v", err)
	}

	// Write todo records
	for _, todo := range tm.todos {
		record := []string{
			strconv.Itoa(todo.ID),
			todo.Description,
			todo.TimeAdded.Format(time.RFC3339),
			todo.TimeCompleted.Format(time.RFC3339),
			strconv.FormatBool(todo.IsCompleted),
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("error writing record: %v", err)
		}
	}

	return nil
}

func NewTodoManager() *TodoManager {
	return &TodoManager{
		nextID: 0,
		todos:  make([]Todo, 0),
	}
}

func (tm *TodoManager) DisplayTodos(all bool) {

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	defer w.Flush()

	if all {
		fmt.Fprintln(w, "ID\tDescription\tCompleted\tTime added\tTime completed")
		for _, t := range tm.todos {
			completedAt := "-"
			if t.IsCompleted {
				completedAt = fmt.Sprintf("%v ago", time.Since(t.TimeCompleted).Round(time.Second))
			}
			fmt.Fprintf(w, "%d\t%s\t%v\t%v ago\t%s\n",
				t.ID,
				t.Description,
				t.IsCompleted,
				time.Since(t.TimeAdded).Round(time.Second),
				completedAt,
			)
		}
		return
	}

	fmt.Fprintln(w, "ID\tDescription")
	for _, t := range tm.todos {
		if !t.IsCompleted {
			fmt.Fprintf(w, "%d\t%s\n", t.ID, t.Description)
		}
	}

}
