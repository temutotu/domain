package service

import (
	"errors"

	"domain/config"
	"domain/domain/repository"
)

const cokkieName string = "session"

type Auth struct {
	repo repository.Repo
}

func NewAuthService() (*Auth, error) {
	repo := repository.GetRepoInterface(config.Conf.Repository)
	if repo == nil {
		return nil, errors.New("failed get repo")
	}

	service := &Auth{
		repo: repo,
	}

	return service, nil
}

func (self *Auth) Authorize(name string, pass string) error {
	result, err := self.repo.Search(name)
	if err != nil {
		return errors.New("search is falied")
	}

	if pass != result {
		return errors.New("authorize is failed")
	}

	return nil
}
