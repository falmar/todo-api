// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func jsonDecode(r io.Reader, into interface{}) error {
	return json.NewDecoder(r).Decode(into)
}

func jsonEncode(w http.ResponseWriter, from interface{}) error {
	return json.NewEncoder(w).Encode(from)
}

func jsonErrorEncode(w http.ResponseWriter, err error, code int, originalError error) {
	log.Println(originalError)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	jsonEncode(w, map[string]map[string]string{
		"message": map[string]string{
			"type":    "ERROR",
			"message": err.Error(),
		},
	})
}
