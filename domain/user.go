package domain

import "errors"

type User struct {
	Name string
	Pass string
}

func CreateUser(name string, pass string) (*User, error) {
	// 簡易validationcheck
	if name == "" || pass == "" {
		return nil, errors.New("name or pass param is empty")
	}

	user := &User{
		Name: name,
		Pass: pass,
	}

	return user, nil

}
