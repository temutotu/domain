package usecase

import (
	"domain/domain"
	"domain/domain/service"
	"errors"
	"net/http"
)

type AuthUsecase struct {
	authService    *service.Auth
	sessionService *service.Session
}

func NewAuthUsecase() (*AuthUsecase, error) {

	authservice, err := service.NewAuthService()
	if err != nil {
		return nil, errors.New("authService fail init")
	}

	sessionservice, err := service.NewSessionService()
	if err != nil {
		return nil, errors.New("sessionService fail init")
	}

	usecase := &AuthUsecase{
		authService:    authservice,
		sessionService: sessionservice,
	}

	return usecase, nil
}

func (self *AuthUsecase) Auth(r *http.Request) *Response {
	response := createResponse()

	//セッションチェック
	if self.sessionService.CheckSession(r) {
		response.Status = 200
		return response
	}
	// 入力データチェック
	name := r.FormValue("name")
	pass := r.FormValue("pass")
	user, err := domain.CreateUser(name, pass)
	if err != nil {
		response.Status = 400
		response.Err = err
		return response
	}
	//認証
	err = self.authService.Authorize(user.Name, user.Pass)
	if err != nil {
		response.Status = 500
		response.Err = err
		return response
	}
	//セッション作成
	cokkie, err := self.sessionService.CreateSession()
	if err != nil {
		response.Status = 500
		response.Err = err
		return response
	}

	response.Status = 200
	response.Err = nil
	response.Cokkie = cokkie
	return response
}
