package storage

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/antoniguss/go-projects/todo-app/types"
)

type CSVStorage struct {
	nextID  int
	csvPath string
	csvFile *os.File
}

var (
	headers = []string{"ID", "Description", "TimeAdded", "TimeCompleted", "IsCompleted"}
)

func NewCSVStorage(csvPath string) (*CSVStorage, error) {
	return &CSVStorage{csvPath: csvPath}, nil
}

func (c *CSVStorage) Setup() error {
	var err error
	c.csvFile, err = loadFile(c.csvPath)
	if err != nil {
		return err
	}

	//Ensure file has correct headers
	scanner := bufio.NewScanner(c.csvFile)
	scanner.Scan()
	presentHeaders := scanner.Text()

	overwrite := presentHeaders == strings.Join(headers, ", ")

	if overwrite {
		c.csvFile.Truncate(0)
		c.csvFile.Seek(0, 0)
		writer := csv.NewWriter(c.csvFile)
		defer writer.Flush()

		if err := writer.Write(headers); err != nil {
			return fmt.Errorf("failed to write headers: %v", err)
		}
	}

	return nil
}

func loadFile(filepath string) (*os.File, error) {
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to open file for reading")
	}

	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		_ = f.Close()
		return nil, err
	}

	return f, nil
}

func closeFile(f *os.File) error {
	syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	return f.Close()
}

func (c *CSVStorage) Close() error {
	return closeFile(c.csvFile)
}

func (c *CSVStorage) AddTodo(description string) error {
	todo := types.Todo{
		ID:          c.nextID,
		Description: description,
		TimeAdded:   time.Now(),
		IsCompleted: false,
	}

	record := []string{
		strconv.Itoa(todo.ID),
		todo.Description,
		todo.TimeAdded.Format(time.RFC3339),
		todo.TimeCompleted.Time.Format(time.RFC3339),
		strconv.FormatBool(todo.IsCompleted),
	}

	writer := csv.NewWriter(c.csvFile)
	defer writer.Flush()

	if err := writer.Write(record); err != nil {
		return fmt.Errorf("failed to write todo to file: %v", err)
	}

	return nil
}

func (c *CSVStorage) GetAllTodos() ([]types.Todo, error) {
	c.csvFile.Seek(0, 0)
	reader := csv.NewReader(c.csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read todos from file: %v", err)
	}

	todos := make([]types.Todo, 0)
	for _, record := range records {
		todo, err := parseTodoRecord(record)
		if err != nil {
			return nil, fmt.Errorf("failed to parse todo: %v", err)
		}
		todos = append(todos, todo)
	}

	return todos, nil
}
func (c *CSVStorage) GetTodo(id int) (types.Todo, error) {
	todos, err := c.GetAllTodos()
	if err != nil {
		return types.Todo{}, fmt.Errorf("failed to get todos: %v", err)
	}

	for _, todo := range todos {
		if todo.ID == id {
			return todo, nil
		}
	}

	return types.Todo{}, fmt.Errorf("todo with id %d not found", id)
}

func (c *CSVStorage) SetCompleted(id int, completed bool) error {
	//TODO implement me
	panic("implement me")
}

func (c *CSVStorage) RemoveTodo(id int) error {
	//TODO implement me
	panic("implement me")
}

//USEFUL CODE?
// if len(tm.todos) == 0 {
// 	return nil
// }

// if !strings.HasSuffix(filePath, ".csv") {
// 	return errors.New("path must be to a .csv file")
// }

// file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
// if err != nil {
// 	return fmt.Errorf("error opening file: %v", err)
// }
// defer file.Close()

// writer := csv.NewWriter(file)
// defer writer.Flush()

// // Write headers
// headers := []string{"ID", "Description", "TimeAdded", "TimeCompleted", "IsCompleted"}
// if err := writer.Write(headers); err != nil {
// 	return fmt.Errorf("error writing headers: %v", err)
// }

// // Write todo records
// for _, todo := range tm.todos {
// 	record := []string{
// 		strconv.Itoa(todo.ID),
// 		todo.Description,
// 		todo.TimeAdded.Format(time.RFC3339),
// 		todo.TimeCompleted.Format(time.RFC3339),
// 		strconv.FormatBool(todo.IsCompleted),
// 	}
// 	if err := writer.Write(record); err != nil {
// 		return fmt.Errorf("error writing record: %v", err)
// 	}
// }

// return nil

// if !strings.HasSuffix(filePath, ".csv") {
// 	return errors.New("path must be to a .csv file")
// }

// file, err := os.Open(filePath)
// if err != nil {
// 	return err
// }
// defer file.Close()

// reader := csv.NewReader(file)

// // Skip header row
// _, err = reader.Read()
// if err != nil {
// 	return fmt.Errorf("error reading header: %v", err)
// }

// var newTodos []Todo
// newTodos = make([]Todo, 0)

// for {
// 	record, err := reader.Read()
// 	if err == io.EOF {
// 		break
// 	}
// 	if err != nil {
// 		return fmt.Errorf("error reading record: %v", err)
// 	}

// 	todo, err := parseTodoRecord(record)
// 	if err != nil {
// 		return fmt.Errorf("error parsing record: %v", err)
// 	}

// 	newTodos = append(newTodos, todo)
// 	if todo.ID >= tm.nextID {
// 		tm.nextID = todo.ID + 1
// 	}
// }

// if overwrite {
// 	tm.todos = make([]Todo, 0)
// 	tm.nextID = 0
// }
// tm.todos = append(tm.todos, newTodos...)

// return nil

func parseTodoRecord(record []string) (types.Todo, error) {
	if len(record) != 5 {
		return types.Todo{}, fmt.Errorf("invalid record length: expected 5, got %d", len(record))
	}

	id, err := strconv.Atoi(record[0])
	if err != nil {
		return types.Todo{}, fmt.Errorf("invalid ID: %v", err)
	}

	timeAdded, err := time.Parse(time.RFC3339, record[2])
	if err != nil {
		return types.Todo{}, fmt.Errorf("invalid TimeAdded: %v", err)
	}

	timeCompleted, err := time.Parse(time.RFC3339, record[3])
	if err != nil {
		return types.Todo{}, fmt.Errorf("invalid TimeCompleted: %v", err)
	}

	isCompleted, err := strconv.ParseBool(record[4])
	if err != nil {
		return types.Todo{}, fmt.Errorf("invalid IsCompleted: %v", err)
	}

	return types.Todo{
		ID:            id,
		Description:   record[1],
		TimeAdded:     timeAdded,
		TimeCompleted: sql.NullTime{Time: timeCompleted, Valid: true},
		IsCompleted:   isCompleted,
	}, nil
}
