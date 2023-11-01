package repository

import (
	"errors"
)

type AssiociativeArray struct {
}

func (assiociativearray *AssiociativeArray) Init() error {
	return nil
}

func (assiociativearray *AssiociativeArray) Search(name string) (string, error) {
	if pass, ok := autharray[name]; ok {
		return pass, nil
	}
	return "", nil
}

func (assiociativearray AssiociativeArray) IsNotDuplicate(name string) bool {
	// nameで検索してヒットしたnameとpassの連想配列
	result, err := assiociativearray.Search(name)
	if err != nil {
		return false
	}

	if result == "" {
		return true
	} else {
		return false
	}
}

func (assiociativearray AssiociativeArray) Add(name string, pass string) error {
	if !assiociativearray.IsNotDuplicate(name) {
		return errors.New("set of name and pass is alrady exist")
	}

	autharray[name] = pass
	return nil
}

func (assiociativearray AssiociativeArray) Delete(name string, pass string) error {
	// nameで検索してヒットしたnameとpassの連想配列
	result, err := assiociativearray.Search(name)
	if err != nil {
		return errors.New("search is failed")
	}

	if pass == result {
		delete(autharray, name)
		return nil
	}

	return errors.New("set of name and pass is not found")
}
