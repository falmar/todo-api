// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

// Todo type that discribes it
type Todo struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Link      string    `json:"link"`
}

func (t Todo) getByID(id int64, db *sql.DB) (*Todo, error) {
	todo := &Todo{}

	ssql := fmt.Sprintf(`SELECT id, user_id, title, completed, created_at, updated_at
		FROM %s.todo t
		WHERE t.id = $1`, os.Getenv("DB_SCHEMA"))

	err := db.QueryRow(ssql, id).Scan(
		&todo.ID,
		&todo.UserID,
		&todo.Title,
		&todo.Completed,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (t Todo) getByUserID(userID int64, db *sql.DB, p *paging) ([]*Todo, error) {
	var count int64

	ssql := fmt.Sprintf(`SELECT COUNT(*) as rCount
		FROM %s.todo t
		WHERE t.user_id = $1`, os.Getenv("DB_SCHEMA"))

	err := db.QueryRow(ssql, userID).Scan(&count)

	if err != nil {
		return nil, err
	}

	p.calc(count)

	ssql = fmt.Sprintf(`SELECT t.id, t.title, t.completed, t.created_at, t.updated_at
		FROM %s.todo t
		WHERE t.user_id = $1 LIMIT $2 OFFSET $3`, os.Getenv("DB_SCHEMA"))

	rows, err := db.Query(ssql, userID, p.Max, p.Init)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	todos := []*Todo{}

	for rows.Next() {
		todo := &Todo{}

		err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)

		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func (t *Todo) insertDB(db *sql.DB) error {
	ssql := fmt.Sprintf(`INSERT INTO %s.todo
		(user_id, title, completed, created_at, updated_at)
		VALUES
		($1, $2, $3, $4, $5) RETURNING id`, os.Getenv("DB_SCHEMA"))

	return db.QueryRow(ssql, t.UserID, t.Title, t.Completed, t.CreatedAt, t.UpdatedAt).Scan(&t.ID)
}

func (t *Todo) updateDB(db *sql.DB) (int64, error) {
	ssql := fmt.Sprintf(`UPDATE %s.todo
		SET title = $1, completed = $2, updated_at = $3
		WHERE id = $4 AND user_id = $5`, os.Getenv("DB_SCHEMA"))

	res, err := db.Exec(ssql, t.Title, t.Completed, t.UpdatedAt, t.ID, t.UserID)

	if err != nil {
		return 0, nil
	}

	return res.RowsAffected()
}
