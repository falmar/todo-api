// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"errors"
	"net/http"
	"time"
)

func registerUserHandler(w http.ResponseWriter, r *http.Request) {
	// check content type is application/json
	if ct := r.Header.Get("Content-Type"); ct != "application/json" {
		err := errors.New("Content-Type: application/json required")
		jsonErrorEncode(w, err, http.StatusBadRequest)
		return
	}

	// declare vars
	response := map[string]interface{}{}
	user := User{}

	// decode request body
	if err := jsonDecode(r.Body, &user); err != nil {
		jsonErrorEncode(w, errJSONMalformed, http.StatusInternalServerError)
		return
	}

	// set current time
	user.CreatedAt = time.Now()

	// validate user is correctly formed for insert
	if err := user.validateDB(postgres); err != nil {
		jsonErrorEncode(w, errJSONMalformed, http.StatusInternalServerError)
		return
	}

	// insert user into db
	if err := user.insertDB(postgres); err != nil {
		jsonErrorEncode(w, errJSONMalformed, http.StatusInternalServerError)
		return
	}

	// set user into response
	response["user"] = user

	// set response header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// write reponse
	if err := jsonEncode(w, response); err != nil {
		jsonErrorEncode(w, errJSONMalformed, http.StatusInternalServerError)
		return
	}
}
