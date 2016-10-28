// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func setRoutes(r *httprouter.Router) {
	common := alice.New(corsMiddleware)
	auth := common.Append(jwtMiddleware)

	r.POST("/login/", wrapperHandler(common.ThenFunc(loginHandler)))
	r.POST("/user/", wrapperHandler(common.ThenFunc(userCreateHandler)))
	r.PUT("/user/:id/", wrapperHandler(auth.ThenFunc(userUpdateHandler)))
}

func wrapperHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), "params", ps)

		h.ServeHTTP(w, r.WithContext(ctx))
	}
}
