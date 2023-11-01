package service

import (
	"fmt"
	"hello/config"
	"hello/domain/repository"
	"io"
	"net/http"
)

type register interface {
	Add() error
	Delete() error
}

func registerService(writer http.ResponseWriter, r *http.Request) {
	enteredName := r.FormValue("name")
	enteredPass := r.FormValue("pass")
	if enteredName == "" || enteredPass == "" {
		io.WriteString(writer, "name or pass param is empty")
		return
	}
	enterdCmd := r.FormValue("cmd")
	if !(enterdCmd == "add" || enterdCmd == "delete") {
		io.WriteString(writer, "cmd is not setting")
		return
	}

	registerobj := RegisterToArray{}
	if enterdCmd == "add" {
		if err := registerobj.Add(enteredName, enteredPass); err != nil {
			io.WriteString(writer, err.Error())
			return
		}

		io.WriteString(writer, "add is success")
	} else if enterdCmd == "delete" {
		if err := registerobj.Delete(enteredName, enteredPass); err != nil {
			io.WriteString(writer, err.Error())
			return
		}

		io.WriteString(writer, "delete is success")
	}
}

func Register(writer http.ResponseWriter, r *http.Request) {
	enteredName := r.FormValue("name")
	enteredPass := r.FormValue("pass")
	if enteredName == "" || enteredPass == "" {
		io.WriteString(writer, "name or pass param is empty")
		return
	}
	repo := repository.GetRepoInterface(config.Conf.Repository)
	err := repo.Add(enteredName, enteredPass)
	fmt.Println(err)
	if err != nil {
		io.WriteString(writer, "register to DB failed")
		return
	}

	io.WriteString(writer, "register is success")
}
