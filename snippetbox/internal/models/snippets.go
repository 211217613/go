package models

import (
	"database/sql"
	"time"
)

// Define a snippet type that models what a data a snippet should hold
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// define a snippetmodel type which wraps a sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

// insert new snippet
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

// return a specific snippet based on id
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(stmt, id)

	s := &Snippet{}

	return nil, nil
}

// return the 10 latest snippets
func (m *SnippetModel) Latest() ([]*SnippetModel, error) {
	return nil, nil
}
