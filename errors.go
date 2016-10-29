// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import "errors"

// bcrypt
var errEncryptPassword = errors.New("Error ocurred trying to encrypt password")

// json
var errMalformedJSON = errors.New("Malformed JSON")
var errJSONContentType = errors.New("Content-Type: application/json required")

// db
var errUniqueConstraintViolationDB = errors.New("Unique Contraint Violation")

// jwt
var errJWTNotFound = errors.New("JWT not found")
