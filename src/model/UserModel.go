package model

import "strconv"

type UserModel struct {
	username string
	password string
	age int
	sex string
}

var UserData map[string]Model

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (u *UserModel) SetUsername(value string) {
	u.username = value
}
func (u *UserModel) SetPassword(value string) {
	u.password = value
}
func (u *UserModel) SetAge(value int) {
	u.age = value
}
func (u *UserModel) SetSex(value string) {
	u.sex = value
}

func (u *UserModel) GetUsername() string {
	return u.username
}
func (u *UserModel) GetPassword() string {
	return u.password
}
func (u *UserModel) GetAge() int {
	return u.age
}
func (u *UserModel) GetSex() string {
	return u.sex
}

// 格式化输出数据信息
func (u *UserModel) ToString() string {
	return u.username + "," + u.password + "," + strconv.Itoa(u.age) + "," + u.sex
}




