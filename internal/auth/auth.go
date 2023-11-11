package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Authorization: ApiKey {orginal APIKEY}
func GetAPIkey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malform auth header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malform in auth header")
	}
	return vals[1], nil
}
