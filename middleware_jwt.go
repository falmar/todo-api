// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"context"
	"net/http"
	"os"
	"strings"
)

func jwtMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var tokenString string
		claims := &MyClaims{}

		if t := r.URL.Query().Get("access_token"); t != "" {
			tokenString = t
		}

		if t := r.Header.Get("Authorization"); t != "" && strings.HasPrefix(t, "Bearer ") {
			if s := strings.Replace(t, "Bearer ", "", 1); s != "" {
				tokenString = strings.TrimSpace(s)
			}
		}

		if tokenString == "" {
			jsonErrorEncode(w, errJWTNotFound, http.StatusBadRequest, errJWTNotFound)
			return
		}

		token, err := parseToken(tokenString, claims, []byte(os.Getenv("JWT_KEY")))

		if err != nil {
			jsonErrorEncode(w, err, http.StatusForbidden, err)
			return
		}

		if token.Valid {
			ctx := context.WithValue(r.Context(), "jwt", token)

			h.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		w.WriteHeader(http.StatusForbidden)

	})
}
