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
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	router := httprouter.New()

	setRoutes(router)

	addr := ":" + os.Getenv("PORT")

	log.Println(fmt.Sprintf("Listening at %s", addr))

	http.ListenAndServe(addr, router)
}
