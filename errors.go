// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"errors"
	"net/http"
)

// bcrypt
var errEncryptPassword = errors.New("Error ocurred trying to encrypt password")

// http
var errInternalServerError = errors.New(http.StatusText(http.StatusInternalServerError))
var errBadRequest = errors.New(http.StatusText(http.StatusBadRequest))
var errNotFound = errors.New(http.StatusText(http.StatusNotFound))
var errUnauthorized = errors.New(http.StatusText(http.StatusUnauthorized))
var errForbidden = errors.New(http.StatusText(http.StatusForbidden))

// json
var errMalformedJSON = errors.New("Malformed JSON")
var errJSONContentType = errors.New("Content-Type: application/json required")

// db
var errUniqueConstraintViolationDB = errors.New("Unique Contraint Violation")

// jwt
var errJWTNotFound = errors.New("JWT not found")
