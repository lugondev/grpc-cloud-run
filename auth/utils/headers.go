package utils

import (
	"net/http"
)

const (
	AuthorizationHeader = "Authorization"
)

func GetAuthorizationHeader(req *http.Request) string {
	return req.Header.Get(AuthorizationHeader)
}

func AddAuthorizationHeaderValue(req *http.Request, value string) {
	req.Header.Add(AuthorizationHeader, value)
}

func DeleteAuthorizationHeaderValue(req *http.Request) {
	req.Header.Del(AuthorizationHeader)
}
