package utils

import (
	"encoding/base64"
	"errors"
	"strings"
)

func DecodeBasicAuth(auth string) (Credentials, error) {
	decoded, err := base64.StdEncoding.DecodeString(auth)
	if err != nil {
		return Credentials{}, err
	}

	parts := strings.Split(string(decoded), ":")
	if len(parts) != 2 {
		return Credentials{}, errors.New("invalid basic auth format")
	}

	return Credentials{Username: parts[0], Password: parts[1]}, nil
}

type Credentials struct {
	Username string
	Password string
}

func Authenticate(username, password string) bool {

	return username == "admin" && password == "password"
}
