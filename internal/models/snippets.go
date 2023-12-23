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

	// Initialize a pointer to a new zeroed Snippet struct
	s := Snippet{}

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

	return &s, nil
}

// returns the 10 most recently created snippets
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`

	// use the query() method on the connection pool to execute our sql statement -> sql.Rows resultset

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	// defer closing to ensure that the resultset is always closed, after checking the error
	// Important! This is critical, as long as a resultset is open, it will keep the underlying db connection open.
	// If something goes wrong, it can lead into a situation where all the connections in db connection pool are being used.
	defer rows.Close()

	snippets := []*Snippet{}

	/*
		Use rows.Next to iterate over the rows in the resultset
		, this prepared the first (and then each subsequent) row to be acted on by the rows.Scan method.
		If iteration over all the rows completes then the result set is automatically closed and the database connection
		is freed.
	*/
	for rows.Next() {
		// create a pointer to a new zeroed Snippet struct
		s := &Snippet{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}
	// Retrieve any error that was encountered during the iteration. Important to check!
	// Don't asume that a successful iteration was completed over the whole resulset
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return snippets, nil

}
