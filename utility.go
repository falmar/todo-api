// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func isJSONContentType(r *http.Request) bool {
	return r.Header.Get("Content-Type") == "application/json"
}

func jsonDecode(r io.Reader, into interface{}) error {
	return json.NewDecoder(r).Decode(into)
}

func jsonEncode(w http.ResponseWriter, from interface{}) error {
	return json.NewEncoder(w).Encode(from)
}

func jsonErrorEncode(w http.ResponseWriter, code int, err error, originalError error) {
	if originalError == nil {
		originalError = errors.New(http.StatusText(code))
	}

	log.Println(originalError)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err == nil {
		err = errors.New(http.StatusText(code))
	}

	jsonEncode(w, map[string]map[string]string{
		"message": map[string]string{
			"type":  "ERROR",
			"title": err.Error(),
		},
	})
}

func paramsFromContext(ctx context.Context, err error) (httprouter.Params, error) {
	if err != nil {
		return nil, err
	}

	if token, ok := ctx.Value("params").(httprouter.Params); ok {
		return token, nil
	}

	return nil, errors.New("type *jwt.Token not found in context")
}
