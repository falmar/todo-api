// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"net/http"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

func userUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if params, ok := r.Context().Value("params").(httprouter.Params); ok {
		if token, ok := r.Context().Value("jwt").(*jwt.Token); ok {
			if claims, ok := token.Claims.(*MyClaims); ok {
				// get UserID from request url
				userID, err := strconv.ParseInt(params.ByName("id"), 10, 64)
				if err != nil || userID == 0 {
					jsonErrorEncode(w, errBadRequest, http.StatusBadRequest, errBadRequest)
					return
				}

				// check if user has permission to update other users or only itself
				canUpdate := isAllowedScope("user:update", claims.Scope) || claims.User.ID == userID
				if !canUpdate {
					jsonErrorEncode(w, errForbidden, http.StatusForbidden, errForbidden)
					return
				}

				// update user

			}
		}
	}

}
