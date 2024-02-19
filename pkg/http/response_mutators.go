package http

import (
	"errors"
	"net/http"
)

type ResponseMutators func(*http.Request) error

func WithTarget(target any) ResponseMutators {
	return func(req *http.Request) error {
		if target == nil {
			return errors.New("you added this option, you might as well provide the target")
		}
		return nil
	}
}
