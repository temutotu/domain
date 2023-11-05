package service

import (
	"domain/config"
	"domain/domain/repository"
	"errors"
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
		return errors.New("register to DB failed")
	}

	return nil
}
