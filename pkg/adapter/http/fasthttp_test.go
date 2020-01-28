package http

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type RequestHandlerMock struct{}

func (rhm RequestHandlerMock) Get(dst []byte, URL string) (statusCode int, body []byte, err error) {
	// todo: throw error and success resp
	if URL == "" {
		return 0, nil, errors.New("connection error")
	}
	return 200, []byte("content"), nil
}

func TestFastHTTP(t *testing.T) {
	sut := fastHTTPClient{
		RequestHandler: &RequestHandlerMock{},
	}

	t.Run("when request is successful", func(t *testing.T) {
		resp, err := sut.Get("http://example.com")

		t.Run("should not have an error", func(t *testing.T) {
			assert.NoError(t, err)
		})

		t.Run("should have a status code of `200`", func(t *testing.T) {
			assert.Equal(t, 200, resp.StatusCode())
		})

		t.Run("should have a content", func(t *testing.T) {
			assert.Equal(t, []byte("content"), resp.Body())
		})
	})
}
