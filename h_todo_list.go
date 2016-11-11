// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"database/sql"
	"errors"
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

	filters := parseTodoFilters(query.Get("filters"), map[string]string{
		"completed": "bool",
	})

	sorts := parseSort(query.Get("sort"), []string{
		"completed",
		"updated_at",
		"created_at",
		"title",
		"id",
	})

	response := map[string]interface{}{}
	todo := Todo{}
	paging := &paging{CurrentPage: page, PageSize: maxPerPage}

	todos, err := todo.getByUserID(postgres, claims.User.ID, filters, sorts, paging)

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

func parseTodoFilters(filterString string, allowedFilters map[string]string) map[string]interface{} {
	foundFilters := map[string]interface{}{}

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

func parseSort(sortString string, allowedSort []string) map[string]string {
	foundSorts := map[string]string{}

	sorts := strings.Split(sortString, ",")

	getByValue := func(s string) (string, error) {
		switch s {
		case "+":
			return "DESC", nil
		case "-":
			return "ASC", nil
		default:
			return "", errors.New("sort value format not valid")
		}
	}

	for _, s := range sorts {
		for _, as := range allowedSort {
			if strings.HasPrefix(s, as) {
				split := strings.SplitN(s, ":", 2)

				if len(split) == 2 {
					if sort, err := getByValue(split[1]); err == nil {
						foundSorts[split[0]] = sort
					}
				}
			}
		}
	}

	return foundSorts
}
