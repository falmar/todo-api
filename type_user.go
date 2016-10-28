// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User struct
type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (u User) getByID(id int64, db *sql.DB) (*User, error) {
	user := &User{}

	ssql := fmt.Sprintf(`
		SELECT u.id, u.name, u.email, u.created_at, u.updated_at
		FROM %s.user u
		WHERE u.id = $1
		`, os.Getenv("DB_SCHEMA"))

	err := db.QueryRow(ssql, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u User) authenticate(e, p string, db *sql.DB) (int64, error) {
	var hash string
	var id int64

	ssql := fmt.Sprintf("SELECT u.id, u.password FROM %s.user u WHERE u.email = $1", os.Getenv("DB_SCHEMA"))

	if err := db.QueryRow(ssql, e).Scan(&id, &hash); err != nil {
		return 0, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(p)); err != nil {
		return 0, err
	}

	return id, nil
}

func (u *User) insertDB(db *sql.DB) error {
	ssql := fmt.Sprintf(`
		INSERT INTO %s.user
		(name, email, password, created_at, updated_at)
		VALUES
		($1, $2, $3, $4, $5) RETURNING id`, os.Getenv("DB_SCHEMA"))

	err := db.QueryRow(ssql, u.Name, u.Email, u.Password, u.CreatedAt, u.UpdatedAt).Scan(&u.ID)

	if err != nil {
		return err
	}

	return nil
}

func (u *User) updateDB(db *sql.DB) error {

	return nil
}

func (u *User) validateDB(db *sql.DB) error {
	var id int64
	var row *sql.Row
	isUpdate := u.ID > 0

	if isUpdate {
		ssql := fmt.Sprintf(`SELECT u.id
								FROM %s.user u
								WHERE u.id != $1 AND
								u.email = $2`, os.Getenv("DB_SCHEMA"))

		row = db.QueryRow(ssql, u.ID, u.Email)
	} else {
		ssql := fmt.Sprintf("SELECT u.id FROM %s.user u WHERE email = $1", os.Getenv("DB_SCHEMA"))
		row = db.QueryRow(ssql, u.Email)
	}

	if err := row.Scan(&id); err != nil && err != sql.ErrNoRows {
		return err
	} else if id > 0 {
		return errUniqueConstraintViolationDB
	}

	return nil
}
