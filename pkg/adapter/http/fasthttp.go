package http

import "cuvva/pkg/adapter/http/fasthttp"

type fastHTTPClient struct {
	fasthttp.RequestHandler
}

func newFastHTTPClient() fastHTTPClient {
	return fastHTTPClient{
		RequestHandler: fasthttp.NewClient(),
	}
}

func (fhc fastHTTPClient) Get(URL string) (ResponseReader, error) {
	statusCode, resp, err := fhc.RequestHandler.Get(nil, URL)
	if err != nil {
		return nil, err
	}

	return &response{
		code:         statusCode,
		responseType: "application/json",
		response:     resp,
	}, nil
}
