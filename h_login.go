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
	if !isJSONContentType(r) {
		jsonErrorEncode(w, http.StatusBadRequest, errJSONContentType, errJSONContentType)
		return
	}

	// decode request
	if err := jsonDecode(r.Body, &request); err != nil {
		jsonErrorEncode(w, http.StatusBadRequest, errMalformedJSON, err)
		return
	}

	// verify email and password are not empty
	if request.Email == "" || request.Password == "" {
		jsonErrorEncode(w, http.StatusBadRequest, nil, nil)
		return
	}

	// user stub
	authUser := User{}

	// authenticate the request
	userID, err := authUser.authenticate(request.Email, request.Password, postgres)
	// verify authentication errors
	if err != nil {
		if err == sql.ErrNoRows {
			jsonErrorEncode(w, http.StatusNotFound, nil, err)
		} else if err == bcrypt.ErrMismatchedHashAndPassword {
			jsonErrorEncode(w, http.StatusUnauthorized, nil, err)
		} else {
			jsonErrorEncode(w, http.StatusInternalServerError, nil, err)
		}
		return
	}

	// get user by id
	user, err := authUser.getByID(userID, postgres)
	if err != nil {
		jsonErrorEncode(w, http.StatusInternalServerError, nil, err)
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
		jsonErrorEncode(w, http.StatusInternalServerError, nil, err)
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
		jsonErrorEncode(w, http.StatusBadRequest, errMalformedJSON, err)
		return
	}
}
