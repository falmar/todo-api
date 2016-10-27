// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func setRoutes(r *httprouter.Router) {

}

func wrapperHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), "params", ps)

		h.ServeHTTP(w, r.WithContext(ctx))
	}
}
