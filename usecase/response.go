package usecase

import "net/http"

type Response struct {
	Status int
	Err    error
	Cokkie *http.Cookie
}

func createResponse() *Response {
	res := &Response{
		Status: 500,
		Err:    nil,
		Cokkie: nil,
	}

	return res
}
