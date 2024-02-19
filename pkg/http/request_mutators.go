package http

import (
	"errors"
	"net/http"
)

type RequestMutators func(*http.Request) error

func WithBearerToken(token string) RequestMutators {
	return func(req *http.Request) error {
		if token == "" {
			return errors.New("you added this option, you might as well provide the token")
		}
		req.Header.Set("Authorization", "Bearer "+token)
		return nil
	}
}
