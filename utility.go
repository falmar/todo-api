// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func isJSONContentType(r *http.Request) bool {
	return r.Header.Get("Content-Type") == "application/json"
}

func jsonDecode(r io.Reader, into interface{}) error {
	return json.NewDecoder(r).Decode(into)
}

func jsonEncode(w http.ResponseWriter, from interface{}) error {
	return json.NewEncoder(w).Encode(from)
}

func jsonErrorEncode(w http.ResponseWriter, code int, err error, originalError error) {
	if originalError == nil {
		originalError = errors.New(http.StatusText(code))
	}

	log.Println(originalError)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err == nil {
		err = errors.New(http.StatusText(code))
	}

	jsonEncode(w, map[string]map[string]string{
		"message": map[string]string{
			"type":  "ERROR",
			"title": err.Error(),
		},
	})
}

func paramsFromContext(ctx context.Context, err error) (httprouter.Params, error) {
	if err != nil {
		return nil, err
	}

	if token, ok := ctx.Value("params").(httprouter.Params); ok {
		return token, nil
	}

	return nil, errors.New("type *jwt.Token not found in context")
}

// taken from: http://stackoverflow.com/a/31551220
func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
