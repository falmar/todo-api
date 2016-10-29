// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"database/sql"
	"net/http"
	"strconv"
)

func todoListHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	token, err := jwtFromContext(ctx, nil)
	claims, err := claimsFromToken(token, err)

	if err != nil {
		jsonErrorEncode(w, http.StatusForbidden, nil, nil)
	}

	page, _ := strconv.ParseInt(r.URL.Query().Get("current_page"), 10, 64)
	maxPerPage, _ := strconv.ParseInt(r.URL.Query().Get("page_size"), 10, 64)

	response := map[string]interface{}{}
	todo := Todo{}
	paging := &paging{CurrentPage: page, PageSize: maxPerPage}

	todos, err := todo.getByUserID(claims.User.ID, postgres, paging)

	if err != nil {
		if err == sql.ErrNoRows {
			jsonErrorEncode(w, http.StatusNotFound, nil, err)
		} else {
			jsonErrorEncode(w, http.StatusInternalServerError, nil, err)
		}
		return
	}

	for _, t := range todos {
		t.Link = r.Host + "/todo/" + strconv.FormatInt(t.ID, 10) + "/"
	}

	pageLinkFormatter := func(n int64) string {
		return r.Host + "/todo/?current_page=" + strconv.FormatInt(n, 10)
	}

	paging.Links = paging.getLinks(pageLinkFormatter)
	paging.PagesURL = paging.getPagesWithURL(pageLinkFormatter)

	response["todos"] = todos
	response["pagination"] = paging

	w.Header().Set("Content-Type", "application/json")

	if err := jsonEncode(w, response); err != nil {
		jsonErrorEncode(w, http.StatusInternalServerError, nil, err)
	}

}
