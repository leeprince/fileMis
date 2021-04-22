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

func (this *AuthService) Register(username, password string, age int, sex string) bool {
	user := model.NewUserModel()
	user.SetUsername(username)
	user.SetPassword(password)
	user.SetSex(sex)
	user.SetAge(age)
	user.Save()
	return true
}
