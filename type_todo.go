// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import "database/sql"

// Todo type that discribes it
type Todo struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"user_id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (t Todo) getByID(id int64, db *sql.DB) (*Todo, error) {
	todo := &Todo{}

	return todo, nil
}

func (t Todo) getByUserID(userID int64, db *sql.DB, limit, offset int64) ([]*Todo, error) {

	return nil, nil
}
