package storage

import (
	"database/sql"
	"fmt"

	"github.com/antoniguss/go-projects/todo-app/manager"
	_ "github.com/mattn/go-sqlite3"
)

var (
	dbPath   = "./todos.db"
	database *sql.DB
)

func Setup() error {
	var err error
	database, err = sql.Open("sqlite3", dbPath)

	if err != nil {
		return err
	}
	sqlStatement := `
  CREATE TABLE IF NOT EXISTS 
  todo (id integer not null primary key, 
        description text);`

	_, err = database.Exec(sqlStatement)
	if err != nil {
		return fmt.Errorf("%q: %s\n", err, sqlStatement)
	}
	return nil
}

func Close() error {
	if database == nil {

		return fmt.Errorf("Database not setup")
	}
	return database.Close()
}

func AddTodo(todo manager.Todo) error {

	if database == nil {
		return fmt.Errorf("Database not setup")
	}

	tx, err := database.Begin()
	if err != nil {
		return fmt.Errorf("Couldn't start transaction: %v\n", err)
	}

	stmt, err := tx.Prepare("INSERT INTO todo (description) VALUES (?);")
	if err != nil {
		return fmt.Errorf("Couldn't create query: %v\n", err)
	}

	_, err = stmt.Exec(todo.Description)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil

}

func GetAllTodos() ([]manager.Todo, error) {

	if database == nil {
		return nil, fmt.Errorf("Database not setup")
	}

	stmt, err := database.Prepare("SELECT * FROM todo;")
	if err != nil {
		return nil, fmt.Errorf("Could not prepare query: %s\n", err)
	}

	queries, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("Could not execute query: %s\n", err)
	}

	todos := make([]manager.Todo, 0)
	for queries.Next() {
		var newTodo manager.Todo

		err := queries.Scan(&newTodo.ID, &newTodo.Description)
		if err != nil {
			return nil, fmt.Errorf("Couldn't parse query row to Todo struct: %s\n", err)
		}

		todos = append(todos, newTodo)

	}

	return todos, nil

}
