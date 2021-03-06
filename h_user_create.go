// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func userCreateHandler(w http.ResponseWriter, r *http.Request) {
	// check content type is application/json
	if !isJSONContentType(r) {
		jsonErrorEncode(w, http.StatusBadRequest, errJSONContentType, errJSONContentType)
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
		jsonErrorEncode(w, http.StatusBadRequest, errMalformedJSON, err)
		return
	}

	if request.Name == "" || request.Email == "" || request.Password == "" {
		jsonErrorEncode(w, http.StatusBadRequest, nil, nil)
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
	np, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		jsonErrorEncode(w, http.StatusInternalServerError, errEncryptPassword, err)
		return
	}
	user.Password = string(np)

	// validate user is correctly formed for insert
	if err := user.validateDB(postgres); err != nil {
		if err == errUniqueConstraintViolationDB {
			jsonErrorEncode(w, http.StatusInternalServerError, errUniqueConstraintViolationDB, err)
		} else {
			jsonErrorEncode(w, http.StatusInternalServerError, nil, err)
		}

		return
	}

	// insert user into db
	if err := user.insertDB(postgres); err != nil {
		jsonErrorEncode(w, http.StatusInternalServerError, nil, err)
		return
	}

	// set user into response
	response["user"] = user
	response["message"] = map[string]string{
		"type":  "SUCCESS",
		"title": "User successfully created",
	}

	// set response header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// write reponse
	if err := jsonEncode(w, response); err != nil {
		jsonErrorEncode(w, http.StatusInternalServerError, errMalformedJSON, err)
		return
	}
}
