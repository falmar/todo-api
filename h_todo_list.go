// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"database/sql"
	"net/http"
)

func todoListHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	token, err := jwtFromContext(ctx, nil)
	claims, err := claimsFromToken(token, err)

	if err != nil {
		jsonErrorEncode(w, http.StatusForbidden, nil, nil)
	}

	response := map[string]interface{}{}

	todo := Todo{}

	todos, err := todo.getByUserID(claims.User.ID, postgres, 15, 0)

	if err != nil {
		if err == sql.ErrNoRows {
			jsonErrorEncode(w, http.StatusNotFound, nil, err)
		} else {
			jsonErrorEncode(w, http.StatusInternalServerError, nil, err)
		}
		return
	}

	response["todos"] = todos

	w.Header().Set("Content-Type", "application/json")

	if err := jsonEncode(w, response); err != nil {
		jsonErrorEncode(w, http.StatusInternalServerError, nil, err)
	}

}
