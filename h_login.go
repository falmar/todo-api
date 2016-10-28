// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"database/sql"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// declare vars
	response := map[string]interface{}{}
	request := struct {
		Email    string `sql:"email"`
		Password string `sql:"password"`
	}{}

	// check request content-type is json
	if ct := r.Header.Get("Content-Type"); ct != "application/json" {
		jsonErrorEncode(w, errJSONContentType, http.StatusBadRequest, errJSONContentType)
	}

	// decode request
	if err := jsonDecode(r.Body, &request); err != nil {
		jsonErrorEncode(w, errMalformedJSON, http.StatusBadRequest, err)
		return
	}

	// verify email and password are not empty
	if request.Email == "" || request.Password == "" {
		jsonErrorEncode(w, errBadRequest, http.StatusBadRequest, errBadRequest)
		return
	}

	// user stub
	authUser := User{}

	// authenticate the request
	userID, err := authUser.authenticate(request.Email, request.Password, postgres)
	// verify authentication errors
	if err != nil {
		if err == sql.ErrNoRows {
			jsonErrorEncode(w, errNotFound, http.StatusNotFound, err)
		} else if err == bcrypt.ErrMismatchedHashAndPassword {
			jsonErrorEncode(w, errUnauthorized, http.StatusUnauthorized, err)
		} else {
			jsonErrorEncode(w, errInternalServerError, http.StatusInternalServerError, err)
		}
		return
	}

	// get user by id
	user, err := authUser.getByID(userID, postgres)
	if err != nil {
		jsonErrorEncode(w, errInternalServerError, http.StatusInternalServerError, err)
		return
	}

	// set claims
	response["claims"] = map[string]interface{}{
		"user":  user,
		"scope": "todo:create,todo:update,todo:delete",
	}
	response["token"] = ""

	w.Header().Set("Content-Type", "application/json")

	if err := jsonEncode(w, response); err != nil {
		jsonErrorEncode(w, errMalformedJSON, http.StatusBadRequest, err)
		return
	}
}
