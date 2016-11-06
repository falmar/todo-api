// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"net/http"
	"os"
)

func checkTokenHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := r.URL.Query().Get("access_token")

	token, err := parseToken(tokenString, &MyClaims{}, []byte(os.Getenv("JWT_KEY")))
	claims, err := claimsFromToken(token, err)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	response := map[string]interface{}{}

	response["claims"] = claims
	response["token"] = token.Raw

	if err := jsonEncode(w, response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}
