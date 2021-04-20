package service

import (
	"errors"
	"fileMis/src/model"
)

type AuthService struct {

}

func (this *AuthService) Login(username, password string) (bool, error) {
	user := model.GetUserByUsername(username)
	if user == nil {
		return false, errors.New("用户名错误")
	}
	if user.GetPassword() != password {
		return false, errors.New("密码错误")
	}
	return true, nil
}
