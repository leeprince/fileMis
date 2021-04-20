package controller

import (
	"fileMis/src/utils"
	"fmt"
)

type AuthController struct {

}

func (this *AuthController) Login() bool  {
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
