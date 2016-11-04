// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"net/http"
	"strconv"
	"time"
)

func todoAddHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	token, err := jwtFromContext(ctx, nil)
	claims, err := claimsFromToken(token, err)

	if err != nil {
		jsonErrorEncode(w, http.StatusForbidden, nil, nil)
		return
	}

	// check content type is application/json
	if !isJSONContentType(r) {
		jsonErrorEncode(w, http.StatusBadRequest, errJSONContentType, errJSONContentType)
		return
	}

	// initial variables
	response := map[string]interface{}{}
	todo := &Todo{}

	// decode response
	if err := jsonDecode(r.Body, todo); err != nil {
		jsonErrorEncode(w, http.StatusBadRequest, nil, err)
		return
	}

	// set default todo variables sanitaze
	todo.ID = 0
	todo.UserID = claims.User.ID
	t := time.Now()
	todo.CreatedAt = t
	todo.UpdatedAt = t

	// insert into db
	if err := todo.insertDB(postgres); err != nil {
		jsonErrorEncode(w, http.StatusInternalServerError, nil, err)
		return
	}

	// generate link
	todo.Link = getHost(r) + "/todo/" + strconv.FormatInt(todo.ID, 10) + "/"

	// create response
	response["todo"] = todo
	response["message"] = map[string]string{
		"type":  "SUCCESS",
		"title": "TODO successfully created",
	}

	// response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// write response
	if err := jsonEncode(w, response); err != nil {
		jsonErrorEncode(w, http.StatusInternalServerError, nil, err)
		return
	}

}
