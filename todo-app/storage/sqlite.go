package storage

import (
	"database/sql"
	"fmt"

	"github.com/antoniguss/go-projects/todo-app/types"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteStorage struct {
	db     *sql.DB
	dbPath string
}

func NewSQLiteStorage(dbPath string) (*SQLiteStorage, error) {
	return &SQLiteStorage{dbPath: dbPath}, nil
}

func (s *SQLiteStorage) Setup() error {
	var err error
	s.db, err = sql.Open("sqlite3", s.dbPath)

	if err != nil {
		return err
	}
	sqlStatement := `
  	CREATE TABLE IF NOT EXISTS todo (
        id INTEGER NOT NULL PRIMARY KEY, 
        description TEXT NOT NULL,
        timeAdded DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        timeCompleted DATETIME,
        isCompleted BOOLEAN DEFAULT FALSE
	);`

	_, err = s.db.Exec(sqlStatement)
	if err != nil {
		return fmt.Errorf("%q: %s\n", err, sqlStatement)
	}
	return nil
}

func (s *SQLiteStorage) Close() error {
	if s.db == nil {
		return fmt.Errorf("Database not setup")
	}
	return s.db.Close()
}

func (s *SQLiteStorage) AddTodo(description string) error {

	if s.db == nil {
		return fmt.Errorf("Database not setup")
	}

	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("Couldn't start transaction: %v\n", err)
	}

	stmt, err := tx.Prepare("INSERT INTO todo (description) VALUES (?);")
	if err != nil {
		return fmt.Errorf("Couldn't create query: %v\n", err)
	}

	_, err = stmt.Exec(description)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil

}

func (s *SQLiteStorage) GetAllTodos() ([]types.Todo, error) {

	if s.db == nil {
		return nil, fmt.Errorf("Database not setup")
	}

	stmt, err := s.db.Prepare("SELECT * FROM todo;")
	if err != nil {
		return nil, fmt.Errorf("Could not prepare query: %s\n", err)
	}

	queries, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("Could not execute query: %s\n", err)
	}

	todos := make([]types.Todo, 0)
	for queries.Next() {
		var newTodo types.Todo

		err := queries.Scan(
			&newTodo.ID,
			&newTodo.Description,
			&newTodo.TimeAdded,
			&newTodo.TimeCompleted,
			&newTodo.IsCompleted,
		)
		if err != nil {
			return nil, fmt.Errorf("Couldn't parse query row to Todo struct: %s\n", err)
		}

		todos = append(todos, newTodo)

	}

	return todos, nil

}

func (s *SQLiteStorage) SetCompleted(id int, completed bool) error {

	if s.db == nil {
		return fmt.Errorf("Database not setup")
	}

	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("Couldn't start transaction: %v\n", err)
	}

	var stmt *sql.Stmt
	if completed {
		stmt, err = tx.Prepare(`
    UPDATE todo 
    SET isCompleted = ?, timeCompleted = datetime('now')
    WHERE id = (?);
    `)

	} else {

		stmt, err = tx.Prepare(`
    UPDATE todo 
    SET isCompleted = ?, timeCompleted = NULL
    WHERE id = (?);
    `)

	}
	if err != nil {
		return fmt.Errorf("Couldn't create query: %v\n", err)
	}

	_, err = stmt.Exec(completed, id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (s *SQLiteStorage) RemoveTodo(id int) error {

	if s.db == nil {
		return fmt.Errorf("Database not setup")
	}

	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("Couldn't start transaction: %v\n", err)
	}

	stmt, err := tx.Prepare(`
    DELETE 
    FROM todo
    WHERE id = (?);
    `)
	if err != nil {
		return fmt.Errorf("Couldn't create query: %v\n", err)
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	if rowsAffected, _ := res.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("Todo with id %d not found", id)
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil

}
