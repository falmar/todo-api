// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"database/sql"
	"net/http"
	"strconv"
)

func todoGetHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	token, err := jwtFromContext(ctx, nil)
	claims, err := claimsFromToken(token, err)
	params, err := paramsFromContext(ctx, err)

	if err != nil {
		jsonErrorEncode(w, http.StatusForbidden, nil, nil)
	}

	todoID, _ := strconv.ParseInt(params.ByName("todo"), 10, 64)

	response := map[string]interface{}{}

	ot := Todo{}

	todo, err := ot.getByID(todoID, postgres)

	if err != nil {
		if err == sql.ErrNoRows {
			jsonErrorEncode(w, http.StatusNotFound, nil, err)
		} else {
			jsonErrorEncode(w, http.StatusInternalServerError, nil, err)
		}
		return
	}

	if todo.UserID != claims.User.ID {
		jsonErrorEncode(w, http.StatusForbidden, nil, nil)
		return
	}

	todo.Link = getHost(r) + "/todo/" + strconv.FormatInt(todo.ID, 10) + "/"

	response["todo"] = todo

	w.Header().Set("Content-Type", "application/json")

	if err := jsonEncode(w, response); err != nil {
		jsonErrorEncode(w, http.StatusInternalServerError, nil, err)
	}

}
