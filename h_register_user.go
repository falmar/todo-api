// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"errors"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func registerUserHandler(w http.ResponseWriter, r *http.Request) {
	// check content type is application/json
	if ct := r.Header.Get("Content-Type"); ct != "application/json" {
		err := errors.New("Content-Type: application/json required")
		jsonErrorEncode(w, err, http.StatusBadRequest, err)
		return
	}

	// declare vars
	response := map[string]interface{}{}
	request := struct {
		User
		Password string `json:"password"`
	}{}

	// decode request body
	if err := jsonDecode(r.Body, &request); err != nil {
		jsonErrorEncode(w, errMalformedJSON, http.StatusInternalServerError, err)
		return
	}

	if request.Name == "" || request.Email == "" || request.Password == "" {
		jsonErrorEncode(w, errBadRequest, http.StatusBadRequest, errBadRequest)
		return
	}

	// set data into user struct
	user := &User{
		Name:      request.Name,
		Email:     request.Email,
		ID:        0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// password encrypt
	if np, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10); err != nil {
		jsonErrorEncode(w, errEncryptPassword, http.StatusInternalServerError, err)
		return
	} else if err == nil {
		user.Password = string(np)
	}

	// validate user is correctly formed for insert
	if err := user.validateDB(postgres); err != nil {
		if err == errUniqueConstraintViolationDB {
			jsonErrorEncode(w, errUniqueConstraintViolationDB, http.StatusInternalServerError, err)
		} else {
			jsonErrorEncode(w, errInternalServerError, http.StatusInternalServerError, err)
		}

		return
	}

	// insert user into db
	if err := user.insertDB(postgres); err != nil {
		jsonErrorEncode(w, errInternalServerError, http.StatusInternalServerError, err)
		return
	}

	// set user into response
	response["user"] = user
	response["message"] = []map[string]string{
		map[string]string{
			"type":    "SUCCESS",
			"message": "User successfully created",
		},
	}

	// set response header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// write reponse
	if err := jsonEncode(w, response); err != nil {
		jsonErrorEncode(w, errMalformedJSON, http.StatusInternalServerError, err)
		return
	}
}
