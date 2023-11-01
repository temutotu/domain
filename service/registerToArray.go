package service

import (
	"hello/domain/repository"
	"net/http"
)

type RegisterToArray struct {
}

func (register RegisterToArray) RegisterService(writer http.ResponseWriter, r *http.Request) {
	registerService(writer, r)
}

func (register RegisterToArray) Add(name string, pass string) error {
	repository := &repository.AssiociativeArray{}
	if err := repository.Add(name, pass); err != nil {
		return err
	}
	return nil
}

func (register RegisterToArray) Delete(name string, pass string) error {
	repository := &repository.AssiociativeArray{}
	if err := repository.Delete(name, pass); err != nil {
		return err
	}
	return nil
}
