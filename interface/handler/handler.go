package handler

import (
	"net/http"
)

func HanlerInit() {
	// ハンドラー関数を定義する
	authhandler := func(w http.ResponseWriter, r *http.Request) {
		Auth(w, r)
	}

	registerhandler := func(w http.ResponseWriter, r *http.Request) {
		Register(w, r)
	}

	// パスとハンドラー関数を結びつける
	http.HandleFunc("/auth/", authhandler)
	http.HandleFunc("/register/", registerhandler)
}
