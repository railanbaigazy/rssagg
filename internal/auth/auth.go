package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(header http.Header) (string, error) {
	const prefix = "ApiKey"

	authString := header.Get("Authorization")
	if authString == "" {
		return "", errors.New("no authentication info found")
	}
	vals := strings.Split(authString, " ")
	if len(vals) != 2 || vals[0] != prefix {
		return "", errors.New("malformed auth header")
	}

	return vals[1], nil
}
