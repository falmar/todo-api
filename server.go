// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// load Environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	addr := ":" + os.Getenv("PORT")

	// set up database connection
	if err := setUpDB(); err != nil {
		log.Fatal(err)
	}

	// set routes
	router := httprouter.New()
	setRoutes(router)

	cors := CORS{router}

	// start listening
	log.Println(fmt.Sprintf("Listening at %s", addr))
	http.ListenAndServe(addr, cors)
}
