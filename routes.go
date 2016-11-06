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
	auth := alice.New(jwtMiddleware)

	// auth
	r.POST("/login/", wrapperHandler(http.HandlerFunc(loginHandler)))

	// user
	r.POST("/user/", wrapperHandler(http.HandlerFunc(userCreateHandler)))
	r.PUT("/user/:id/", wrapperHandler(auth.ThenFunc(userUpdateHandler)))
	r.DELETE("/user/:id/", wrapperHandler(auth.ThenFunc(userDeleteHandler)))

	// todo
	r.GET("/todo/", wrapperHandler(auth.ThenFunc(todoListHandler)))
	r.POST("/todo/", wrapperHandler(auth.ThenFunc(todoAddHandler)))
	r.PUT("/todo/:todo/", wrapperHandler(auth.ThenFunc(todoUpdateHandler)))
	r.DELETE("/todo/:todo/", wrapperHandler(auth.ThenFunc(todoDeleteHandler)))
}

func wrapperHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), "params", ps)

		h.ServeHTTP(w, r.WithContext(ctx))
	}
}
