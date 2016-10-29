// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"net/http"
	"strconv"
)

func userDeleteHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params, err := paramsFromContext(ctx, nil)
	token, err := jwtFromContext(ctx, err)
	claims, err := claimsFromToken(token, err)

	if err != nil {
		jsonErrorEncode(w, errForbidden, http.StatusForbidden, errForbidden)
	}

	// get UserID from request url
	userID, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || userID == 0 {
		jsonErrorEncode(w, errBadRequest, http.StatusBadRequest, errBadRequest)
		return
	}

	// check if user has permission to update other users or only itself
	canDelete := isAllowedScope("user:delete", claims.Scope) && claims.User.ID != userID
	if !canDelete {
		jsonErrorEncode(w, errForbidden, http.StatusForbidden, errForbidden)
		return
	}

	// start update user
	response := map[string]interface{}{}

	// set data into user struct
	user := &User{ID: userID}

	// update user
	if rows, err := user.deleteDB(postgres); err != nil {
		jsonErrorEncode(w, errInternalServerError, http.StatusInternalServerError, err)
		return
	} else if rows == 0 {
		jsonErrorEncode(w, errNotFound, http.StatusNotFound, errNotFound)
		return
	}

	// set user into response
	response["message"] = map[string]string{
		"type":  "SUCCESS",
		"title": "User successfully deleted",
	}

	// set response header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// write reponse
	if err := jsonEncode(w, response); err != nil {
		jsonErrorEncode(w, errMalformedJSON, http.StatusInternalServerError, err)
		return
	}
}
