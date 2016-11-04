// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"net/http"
	"strconv"
	"time"
)

func todoUpdateHandler(w http.ResponseWriter, r *http.Request) {
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

	// check content type is application/json
	if !isJSONContentType(r) {
		jsonErrorEncode(w, http.StatusBadRequest, errJSONContentType, errJSONContentType)
		return
	}

	// initial variables
	response := map[string]interface{}{}
	todo := &Todo{}

	// decode response
	if err = jsonDecode(r.Body, todo); err != nil {
		jsonErrorEncode(w, http.StatusBadRequest, nil, err)
		return
	}

	// set default todo variables sanitaze
	todo.ID = TodoID
	todo.UserID = claims.User.ID
	todo.UpdatedAt = time.Now()

	// insert into db
	rows, err := todo.updateDB(postgres)
	if err != nil {
		jsonErrorEncode(w, http.StatusInternalServerError, nil, err)
		return
	} else if rows == 0 {
		jsonErrorEncode(w, http.StatusNotFound, nil, nil)
		return
	}

	uTodo, err := todo.getByID(todo.ID, postgres)

	if err != nil {
		jsonErrorEncode(w, http.StatusInternalServerError, nil, err)
		return
	}

	// generate link
	uTodo.Link = getHost(r) + "/todo/" + strconv.FormatInt(uTodo.ID, 10) + "/"

	// create response
	response["todo"] = uTodo
	response["message"] = map[string]string{
		"type":  "SUCCESS",
		"title": "TODO successfully updated",
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
