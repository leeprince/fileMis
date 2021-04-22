package controller

import "fileMis/src/service"

var (
	Next string = indexControllerMark+"::Welcome" // 默认路由
	View string
)

var (
	Views       map[string][][3]string
	Controllers map[string]interface{}
	authService *service.AuthService
	userService *service.UserService
)

// 注册的控制器标识
var (
	indexControllerMark string = "IndexController"
	userControllerMark string = "UserController"
	authControllerMark string = "AuthController"
)

func init() {
	Views = make(map[string][][3]string)
	Controllers = make(map[string]interface{})
	
	// 添加控制器方法
	initView()
	registerController()
}

func registerController()  {
	Controllers[authControllerMark] = &AuthController{}
	Controllers[indexControllerMark] = &IndexController{}
	Controllers[userControllerMark] = &UserController{}
}

func initView() {
	Views["loginView"] = [][3]string{
		{
			authControllerMark,
			"Login",
			"登录系统",
		},
		{
			authControllerMark,
			"Register",
			"注册用户",
		},
	}
	Views["indexView"] = [][3]string{
		{
			indexControllerMark,
			"Index",
			"进入首页",
		},
		{
			userControllerMark,
			"List",
			"展示用户信息",
		},
	}
	
}
