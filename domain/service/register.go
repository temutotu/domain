package service

import (
	"errors"
	"fmt"
	"hello/config"
	"hello/domain/repository"
)

type Register struct {
	repo repository.Repo
}

func NewRegisterService() (*Register, error) {
	repo := repository.GetRepoInterface(config.Conf.Repository)
	if repo == nil {
		return nil, errors.New("failed get repo")
	}

	service := &Register{
		repo: repo,
	}

	return service, nil
}

func (self *Register) Register(name string, pass string) error {
	err := self.repo.Add(name, pass)
	if err != nil {
		fmt.Println(err)
		return errors.New("register to DB failed")
	}

	return nil
}
