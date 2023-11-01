package service

import (
	"fmt"
	"io"
	"net/http"

	"hello/domain/config"
	"hello/domain/repository"
	"hello/domain/server/session"
)

const cokkieName string = "session"

func Authorize(writer http.ResponseWriter, r *http.Request) {
	// 最初にセッションがあるかチェック
	_, err := r.Cookie(cokkieName)
	if err == nil {
		http.Redirect(writer, r, "http://localhost:8080/main/", 301)
		return
	}
	enteredName := r.FormValue("name")
	enteredPass := r.FormValue("pass")
	if enteredName == "" || enteredPass == "" {
		io.WriteString(writer, "name or pass param is empty")
		return
	}

	//認証処理
	repo := repository.GetRepoInterface(config.Conf.Repository)
	result, err := repo.Search(enteredName)
	if err != nil {
		fmt.Println(err)
		io.WriteString(writer, "search is falied")
		return
	}

	if enteredPass != result {
		io.WriteString(writer, "authorize is failed")
		return
	}
	//認証処理終わり

	//　ここでセッションを作成
	sessionID, err := session.Create()
	if err != nil {
		fmt.Println(err)
		io.WriteString(writer, "session create is failed")
		return
	}

	cookie := &http.Cookie{
		Name:  "session",
		Value: sessionID,
	}
	http.SetCookie(writer, cookie)
	io.WriteString(writer, "authorize is successed")
}

func AuthorizeFromArray(writer http.ResponseWriter, r *http.Request) {
	// 最初にセッションがあるかチェック
	_, err := r.Cookie(cokkieName)
	if err == nil {
		http.Redirect(writer, r, "http://localhost:8080/main/", 301)
		return
	}
	enteredName := r.FormValue("name")
	enteredPass := r.FormValue("pass")
	if enteredName == "" || enteredPass == "" {
		io.WriteString(writer, "name or pass param is empty")
		return
	}

	repository := &repository.AssiociativeArray{}
	result, _ := repository.Search(enteredName)
	if enteredPass == result {
		io.WriteString(writer, "authorize is successed")
		return
	}

	io.WriteString(writer, "authorize is faied")
}

func AuthorizeFromDB(writer http.ResponseWriter, r *http.Request) {
	// 最初にセッションがあるかチェック
	_, err := r.Cookie(cokkieName)
	if err == nil {
		http.Redirect(writer, r, "http://localhost:8080/main/", 301)
		return
	}
	enteredName := r.FormValue("name")
	enteredPass := r.FormValue("pass")
	if enteredName == "" || enteredPass == "" {
		io.WriteString(writer, "name or pass param is empty")
		return
	}
	// result is repository.User struct
	repo := repository.GetRepoInterface("MySQL")
	result, err := repo.Search(enteredName)
	if err != nil {
		fmt.Println(err)
		io.WriteString(writer, "search is falied")
		return
	}

	if enteredPass == result {
		io.WriteString(writer, "authorize is successed")
		return
	}

	io.WriteString(writer, "authorize is faied")

}
