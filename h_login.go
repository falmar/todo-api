// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"database/sql"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
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
	myClaims := MyClaims{
		User:  *user,
		Scope: "todo:create,todo:update,todo:delete",
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			Issuer:    os.Getenv("JWT_ISSUER"),
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}

	token, err := generateToken(myClaims, []byte(os.Getenv("JWT_KEY")))

	if err != nil {
		jsonErrorEncode(w, errInternalServerError, http.StatusInternalServerError, err)
		return
	}

	response["claims"] = myClaims
	response["token"] = token
	response["message"] = map[string]string{
		"title": "User successfully log in",
		"type":  "SUCCESS",
	}

	w.Header().Set("Content-Type", "application/json")

	if err := jsonEncode(w, response); err != nil {
		jsonErrorEncode(w, errMalformedJSON, http.StatusBadRequest, err)
		return
	}
}
