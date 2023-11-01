package service

import (
	"io"
	"net/http"
)

func MainManue(writer http.ResponseWriter, r *http.Request) {
	io.WriteString(writer, "here is mainmanue")
}
