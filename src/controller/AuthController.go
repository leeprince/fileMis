package controller

import (
	"fileMis/src/utils"
	"fmt"
	"strconv"
)

type AuthController struct {
}

func (this *AuthController) Login() bool {
	fmt.Print("输入您的用户名：")
	userName := utils.CRead()
	fmt.Print("输入您的密码：")
	password := utils.CRead()
	
	fmt.Println("用户名：", userName)
	fmt.Println("密码：", password)
	
	if bool, err := authService.Login(userName, password); ! bool {
		fmt.Println("登录失败:", err)
		View = "loginView"
		return false
	}
	fmt.Println("登录成功")
	
	View = "indexView"
	return true
}

func (this *AuthController) Register() bool {
	fmt.Println("输入你需要注册的用户信息 username,password,age,sex")
	fmt.Print("输入你的用户名 : ")
	username := utils.CRead()
	fmt.Print("输入你的密码 : ")
	password := utils.CRead()
	fmt.Print("确认密码 : ")
	password1 := utils.CRead()
	fmt.Print("输入你的年龄 : ")
	age, _ := strconv.Atoi(utils.CRead())
	fmt.Print("输入你的性别 : ")
	sex := utils.CRead()
	
	if password != password1 {
		fmt.Println("注册失败：密码需要一致")
		return false
	}
	
	if ! authService.Register(username, password, age, sex) {
		fmt.Println("注册失败")
		return false
	}
	
	fmt.Println("注册成功")
	
	View = "loginView"
	return true
}
