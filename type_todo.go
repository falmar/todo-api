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
