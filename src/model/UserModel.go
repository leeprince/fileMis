package model

import (
	"strconv"
)

type UserModel struct {
	username string
	password string
	age int
	sex string
}

var (
	userModelTable string = "user" // 用户信息
	userData map[string]Model // 用户信息
	userDataKey string = "username" // 存储用户信息的主键
)

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
func (u *UserModel) All() []*UserModel  {
	var users []*UserModel = make([]*UserModel, 0)
	for _, user := range userData {
		users = append(users, user.(*UserModel))
	}
	return users
}

// 格式化输出数据信息
func (u *UserModel) ToString() string {
	return u.username + "," + u.password + "," + strconv.Itoa(u.age) + "," + u.sex
}

func GetUserByUsername(username string) *UserModel {
	// fmt.Println("GetUser by username:", username)
	// fmt.Printf("userData 类型：%T 值：%v \n", userData[username], userData[username])
	
	data, ok := userData[username]
	if !ok {
		return nil
	}
	return data.(*UserModel)
}

func (u *UserModel) Save() bool {
	// 先在内存里面保存一份
	userData[u.username] = u
	return fwrite(userModelTable, userData)
}




