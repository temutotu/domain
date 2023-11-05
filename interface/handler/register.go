package handler

import (
	"hello/usecase"
	"io"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {

	usecase, err := usecase.NewRegisterUsecase()
	if err != nil {
		w.WriteHeader(500)
		io.WriteString(w, err.Error())
		return
	}

	response := usecase.Register(r)
	w.WriteHeader(response.Status)
	if response.Err == nil {
		if response.Cokkie != nil {
			http.SetCookie(w, response.Cokkie)
		}
		io.WriteString(w, "register is succsess")
	} else {
		io.WriteString(w, response.Err.Error())
	}
	return
}
