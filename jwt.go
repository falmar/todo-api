// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// MyClaims custom claim to use in this app
type MyClaims struct {
	jwt.StandardClaims
	User  User   `json:"user"`
	Scope string `json:"scope"`
}

// Valid implement interface jwt.Claims
func (mc MyClaims) Valid() error {
	if mc.User.ID == 0 {
		return errors.New("User's ID can not be empty")
	}

	if mc.Scope == "" {
		return errors.New("Scope can not be empty")
	}

	if mc.ExpiresAt < time.Now().Unix() {
		return errors.New("JWT expired")
	}

	return nil
}

func generateToken(claims jwt.Claims, secretKey []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

func parseToken(tokenString string, claims jwt.Claims, secretKey []byte) (*jwt.Token, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		// validate the alg
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	}

	return jwt.ParseWithClaims(tokenString, claims, keyFunc)
}

func isAllowedScope(s, ss string) bool {
	scopeSlice := strings.Split(ss, ",")

	for _, rs := range scopeSlice {
		if rs == s {
			return true
		}
	}

	return false
}

func jwtFromContext(ctx context.Context, err error) (*jwt.Token, error) {
	if err != nil {
		return nil, err
	}

	if token, ok := ctx.Value("jwt").(*jwt.Token); ok {
		return token, nil
	}

	return nil, errors.New("type *jwt.Token not found in context")
}

func claimsFromToken(t *jwt.Token, err error) (*MyClaims, error) {
	if err != nil {
		return nil, err
	}

	if claims, ok := t.Claims.(*MyClaims); ok {
		return claims, nil
	}

	return nil, errors.New("type *MyClaims not found in *jwt.Token")
}
