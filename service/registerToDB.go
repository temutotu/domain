package service

import (
	"hello/domain/repository"
	"net/http"
)

type RegisterToDB struct {
}

func (register RegisterToDB) RegisterService(writer http.ResponseWriter, r *http.Request) {
	registerService(writer, r)
}

func (register RegisterToDB) Add(name string, pass string) error {
	// validationcneck

	if err := repository.InsertMySQL(name, pass); err != nil {
		return err
	}

	return nil
}

func (register RegisterToDB) Delete(name string, pass string) error {
	return nil
}
