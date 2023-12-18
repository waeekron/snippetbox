package models

import (
	"database/sql"
	"errors"
	"time"
)

// snippet struct to hold the data of an individual struct.
// corresponds to the fields in mysql table snippets
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// SnippetModel wraps *sql.DB connection pool, it will access and modify DB
type SnippetModel struct {
	DB *sql.DB
}

// inserts a new snippet to the database
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title,content, created, expires)
	VALUES(?,?,UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP, INTERVAL ? DAY))`

	// Execute the statement with given params
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return int(id), nil
}

// returns snippet by id
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	// sql statement we want to execute, ? parameter
	stmt := `SELECT id, title, content, created ,expires FROM snippets WHERE expires > UTC_TIMEsTAMP() AND id = ?`

	row := m.DB.QueryRow(stmt, id)

	// Initialize a pointer to a new zeroed Snipper struct
	s := &Snippet{}

	// Use row.Scan() to copy the values from each field in sql.Row to the corresponding field in the Snippet struct.

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		// if the query returns no rows, Scan will return sql.ErrNoRows
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}
	// If all is OK, we return the Snippet struct

	return s, nil
}

// returns the 10 most recently created snippets
func (m *SnippetModel) Latest() (*[]Snippet, error) {
	return nil, nil
}
