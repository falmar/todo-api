// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"database/sql"
	"time"
)

// User struct
type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) insertDB(db *sql.DB) error {

	return nil
}

func (u *User) updateDB(db *sql.DB) error {

	return nil
}

func (u *User) validateDB(db *sql.DB) error {

	return nil
}
