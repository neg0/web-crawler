package http

type ResponseReader interface {
	StatusCode() int
	ContentType() string
	Body() []byte
}

type response struct {
	code         int
	responseType string
	response     []byte
}

func (r response) StatusCode() int {
	return r.code
}

func (r response) ContentType() string {
	return r.responseType
}

func (r response) Body() []byte {
	return r.response
}
