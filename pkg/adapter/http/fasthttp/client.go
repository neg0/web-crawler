package fasthttp

import (
	"github.com/valyala/fasthttp"
)

type RequestHandler interface {
	Get(dst []byte, url string) (statusCode int, body []byte, err error)
}

type Client struct {
	RequestHandler
}

func NewClient() Client {
	return Client{
		RequestHandler: &fasthttp.Client{},
	}
}

func (c Client) Get(dst []byte, URL string) (statusCode int, body []byte, err error) {
	return c.RequestHandler.Get(nil, URL)
}
