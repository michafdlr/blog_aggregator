package main

import (
	"errors"
	"log"
	"net/http"
	"strings"
)

func GetHeader(header http.Header) (string, error) {
	tokenString, found := strings.CutPrefix(header.Get("Authorization"), "ApiKey ")
	log.Printf("Token:%s", tokenString)
	if !found {
		return "", errors.New("couldn't find ApiKey")
	}
	return tokenString, nil
}
