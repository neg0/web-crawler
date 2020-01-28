package http

import (
	"fmt"
	"net/http"
)

var acceptedStatusCodes = []int{
	http.StatusOK,
	http.StatusAccepted,
}

type RequestHandler interface {
	Get(URL string) (ResponseReader, error)
}

type Client struct {
	RequestHandler
}

// Creates a new instance of HTTP Client to be used in the Core Domains
func NewClient() Client {
	return Client{
		RequestHandler: newFastHTTPClient(),
	}
}

// Creates GET HTTP request to parameter set as `URL`
func (c Client) Get(URL string) (ResponseReader, error) {
	resp, err := c.RequestHandler.Get(URL)
	if err != nil {
		return nil, err
	}

	hasError := hasStatusCodeError(resp)
	if hasError != nil {
		return nil, hasError
	}

	return resp, nil
}

func hasStatusCodeError(resp ResponseReader) error {
	for _, statusCode := range acceptedStatusCodes {
		if statusCode == resp.StatusCode() {
			return nil
		}
	}

	return fmt.Errorf("status code is: %d. Error response is: %s", resp.StatusCode(), resp.Body())
}
