package usecase

import (
	"domain/domain"
	"domain/domain/service"
	"errors"
	"net/http"
)

type RegisterUsecase struct {
	registerService *service.Register
}

func NewRegisterUsecase() (*RegisterUsecase, error) {

	registerservice, err := service.NewRegisterService()
	if err != nil {
		return nil, errors.New("authService fail init")
	}

	usecase := &RegisterUsecase{
		registerService: registerservice,
	}

	return usecase, nil
}

func (self *RegisterUsecase) Register(r *http.Request) *Response {
	response := createResponse()

	name := r.FormValue("name")
	pass := r.FormValue("pass")
	user, err := domain.CreateUser(name, pass)
	if err != nil {
		response.Status = 400
		response.Err = err
		return response
	}

	err = self.registerService.Register(user.Name, user.Pass)
	if err != nil {
		response.Status = 500
		response.Err = err
		return response
	}

	response.Status = 200
	response.Err = err
	return response
}
