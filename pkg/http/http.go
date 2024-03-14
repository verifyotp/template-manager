package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type RequestOptions struct {
	URL        string
	Method     string
	httpClient *http.Client
	Body       map[string]interface{}
	Target     any
}

type Requester interface {
	MakeRequest(req *RequestOptions) error
}

func MakeRequest(ctx context.Context, options *RequestOptions) error {
	if options.URL == "" {
		return errors.New("missing URL parameter")
	}
	if options.Method == "" {
		return errors.New("missing Method parameter")
	}

	requestBody, err := json.Marshal(options.Body)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(options.Method, options.URL, bytes.NewReader(requestBody))
	if err != nil {
		return err
	}

	resp, err := options.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) { _ = Body.Close() }(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	if options.Target != nil {
		err = json.Unmarshal(body, options.Target)
		if err != nil {
			return err

		}
	}

	return nil
}
