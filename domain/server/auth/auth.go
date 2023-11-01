package authmanager

import (
	"errors"
	"strings"
)

// ユーザ名とパスワードのテーブル　デバッグ用
var authtable = map[string]string{
	"taro": "japan",
	"mike": "america",
	"leon": "german",
}

func CheckNameAndPass(name string, pass string) error {
	if val, ok := authtable[name]; ok {
		if strings.Compare(val, pass) == 0 {
			return nil
		}
	}
	return errors.New("name or pass is disagreement")
}
