// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

var postgres *sql.DB

func setUpDB() error {
	var err error

	connString := `
		host=` + os.Getenv("DB_HOST") + `
		user=` + os.Getenv("DB_USER") + `
		dbname=` + os.Getenv("DB_NAME") + `
		password=` + os.Getenv("DB_PASSWORD") + `
		sslmode=` + os.Getenv("DB_SSLMODE")

	postgres, err = sql.Open("postgres", connString)

	if err != nil {
		return err
	}

	return postgres.Ping()
}
