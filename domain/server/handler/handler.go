package handler

import (
	"hello/service"
	"net/http"
)

type handler interface {
	commonHanlerInit() //共通のハンドラを初期化する関数
	extraHandlerInit() //クラス固有のハンドラを初期化する関数
}

func defaultCommonHanlerInit() {
	// ハンドラー関数を定義する
	authhandler := func(w http.ResponseWriter, r *http.Request) {
		//service.Authorize(w, r)
		service.Authorize(w, r)
	}

	registerhandler := func(w http.ResponseWriter, r *http.Request) {
		service.Register(w, r)
	}

	mainmanuehandler := func(w http.ResponseWriter, r *http.Request) {
		service.MainManue(w, r)
	}

	// パスとハンドラー関数を結びつける
	http.HandleFunc("/auth/", authhandler)
	http.HandleFunc("/main/", mainmanuehandler)
	http.HandleFunc("/register/", registerhandler)
}
