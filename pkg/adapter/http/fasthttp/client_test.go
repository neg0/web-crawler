package fasthttp

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type RequestHandlerMock struct{}

func (r RequestHandlerMock) Get(dst []byte, URL string) (int, []byte, error) {
	if URL == "" {
		return 0, nil, errors.New("connection error")
	}
	return 200, []byte("response body content"), nil
}

func TestClient(t *testing.T) {
	var sut Client

	t.Run("when request is successful", func(t *testing.T) {
		sut = Client{
			&RequestHandlerMock{},
		}
		statusCode, resp, err := sut.Get(nil, "http://example.com/about")

		t.Run("should not have an error", func(t *testing.T) {
			assert.NoError(t, err)
		})

		t.Run("should have a status code of `200`", func(t *testing.T) {
			assert.Equal(t, 200, statusCode)
		})

		t.Run("should have a response body is `byte` type", func(t *testing.T) {
			assert.IsType(t, []byte{}, resp)
		})
	})

	t.Run("when request is unsuccessful", func(t *testing.T) {
		sut = Client{
			&RequestHandlerMock{},
		}
		_, resp, err := sut.Get(nil, "")

		t.Run("should not have an error", func(t *testing.T) {
			assert.Error(t, err)
		})

		t.Run("should have an empty response body", func(t *testing.T) {
			assert.Len(t, resp, 0)
		})
	})
}
