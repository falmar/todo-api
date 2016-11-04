// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"net/http"
	"strconv"
)

func todoDeleteHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params, err := paramsFromContext(ctx, nil)
	token, err := jwtFromContext(ctx, err)
	claims, err := claimsFromToken(token, err)

	if err != nil {
		jsonErrorEncode(w, http.StatusForbidden, nil, nil)
		return
	}

	// get todo ID from params
	TodoID, err := strconv.ParseInt(params.ByName("todo"), 10, 64)
	if err != nil || TodoID == 0 {
		jsonErrorEncode(w, http.StatusBadRequest, nil, err)
		return
	}

	// initial variables
	response := map[string]interface{}{}
	todo := &Todo{ID: TodoID, UserID: claims.User.ID}

	// insert into db
	rows, err := todo.deleteDB(postgres)
	if err != nil {
		jsonErrorEncode(w, http.StatusInternalServerError, nil, err)
		return
	} else if rows == 0 {
		jsonErrorEncode(w, http.StatusNotFound, nil, nil)
		return
	}

	// create response
	response["message"] = map[string]string{
		"type":  "SUCCESS",
		"title": "TODO successfully deleted",
	}

	// response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// write response
	if err := jsonEncode(w, response); err != nil {
		jsonErrorEncode(w, http.StatusInternalServerError, nil, err)
		return
	}

}
