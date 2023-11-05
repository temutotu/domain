package test

import (
	"net/http"
	"testing"
)

func TestAuth(t *testing.T) {
	//認証成功
	url := "http://localhost:8080/auth?name=taro&pass=japan"
	response, err := http.Get(url)
	defer response.Body.Close()
	if response.StatusCode != 200 {
		t.Errorf("authsuccess test faied status:%d err:%s", response.StatusCode, err)
	}
	//認証失敗
	url = "http://localhost:8080/auth?name=taro&pass=nihon"
	response, err = http.Get(url)
	if response.StatusCode != 500 {
		t.Errorf("authfail test1 faied status:%d err:%s", response.StatusCode, err)
	}

	url = "http://localhost:8080/auth?name=jiro&pass=japan"
	response, err = http.Get(url)
	if response.StatusCode != 500 {
		t.Errorf("authfail test2 faied status:%d err:%s", response.StatusCode, err)
	}
	//validationerr
	url = "http://localhost:8080/auth?name=taro&pass="
	response, err = http.Get(url)
	if response.StatusCode != 400 {
		t.Errorf("authvalidationerr test1 faied status:%d err:%s", response.StatusCode, err)
	}

	url = "http://localhost:8080/auth?name=&pass=japan"
	response, err = http.Get(url)
	if response.StatusCode != 400 {
		t.Errorf("authvalidationerr test1 faied status:%d err:%s", response.StatusCode, err)
	}
}

func TestRegister(t *testing.T) {
	url := "http://localhost:8080/register/?name=test&pass=test"
	response, err := http.Get(url)
	defer response.Body.Close()
	if response.StatusCode != 200 {
		t.Errorf("registersuccess test faied status:%d err:%s", response.StatusCode, err)
	}

	url = "http://localhost:8080/register/?name=taro&pass=japan"
	response, err = http.Get(url)
	if response.StatusCode != 500 {
		t.Errorf("registerfail test1 faied status:%d err:%d", response.StatusCode, err)
	}

	url = "http://localhost:8080/register/?name=test&pass=japan"
	response, err = http.Get(url)
	if response.StatusCode != 500 {
		t.Errorf("registerfail test2 faied status:%d err:%d", response.StatusCode, err)
	}

	url = "http://localhost:8080/register/?name=test2&pass="
	response, err = http.Get(url)
	if response.StatusCode != 400 {
		t.Errorf("validationtest test1 faied status:%d err:%s", response.StatusCode, err)
	}

	url = "http://localhost:8080/register/?name=&pass=test2"
	response, err = http.Get(url)
	if response.StatusCode != 400 {
		t.Errorf("validationtest test1 faied status:%d err:%s", response.StatusCode, err)
	}
}
