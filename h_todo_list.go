// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"
)

func todoListHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	token, err := jwtFromContext(ctx, nil)
	claims, err := claimsFromToken(token, err)

	if err != nil {
		jsonErrorEncode(w, http.StatusForbidden, nil, nil)
	}

	query := r.URL.Query()

	page, _ := strconv.ParseInt(query.Get("current_page"), 10, 64)
	maxPerPage, _ := strconv.ParseInt(query.Get("page_size"), 10, 64)

	filters := parseTodoFilters(query.Get("filters"))

	response := map[string]interface{}{}
	todo := Todo{}
	paging := &paging{CurrentPage: page, PageSize: maxPerPage}

	todos, err := todo.getByUserID(postgres, claims.User.ID, filters, paging)

	if err != nil {
		if err == sql.ErrNoRows {
			jsonErrorEncode(w, http.StatusNotFound, nil, err)
		} else {
			jsonErrorEncode(w, http.StatusInternalServerError, nil, err)
		}
		return
	}

	for _, t := range todos {
		t.Link = getHost(r) + "/todo/" + strconv.FormatInt(t.ID, 10) + "/"
	}

	pageLinkFormatter := func(n int64) string {
		return getHost(r) + "/todo/?current_page=" + strconv.FormatInt(n, 10)
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

func parseTodoFilters(filterString string) map[string]interface{} {
	foundFilters := map[string]interface{}{}
	allowedFilters := map[string]string{
		"completed": "bool",
	}

	filters := strings.Split(filterString, ",")

	getByType := func(str, t string) (interface{}, error) {
		switch t {
		case "bool":
			return strconv.ParseBool(str)
		default:
			return str, nil
		}

	}

	// loop splitted filters if any
	for _, filter := range filters {

		// loop allowed filters
		for af, aft := range allowedFilters {

			// check if is allowed
			if strings.HasPrefix(filter, af) {
				split := strings.SplitN(filter, ":", 2)

				if len(split) == 2 {
					if t, err := getByType(split[1], aft); err == nil {
						foundFilters[af] = t
					}
				}

			}
		}

	}

	return foundFilters
}
