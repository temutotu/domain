package handler

import (
	"hello/usecase"
	"io"
	"net/http"
)

func Auth(w http.ResponseWriter, r *http.Request) {

	usecase, err := usecase.NewAuthUsecase()
	if err != nil {
		w.WriteHeader(500)
		io.WriteString(w, err.Error())
		return
	}

	response := usecase.Auth(r)
	w.WriteHeader(response.Status)
	if response.Err == nil {
		io.WriteString(w, "auth is succsess")
	} else {
		io.WriteString(w, response.Err.Error())
	}
	return
}
