package controller

import "fmt"

type UserController struct {
}

func (this *UserController) List() {
	users := userService.GetList()
	fmt.Println("| username  | password | age | sex |")
	for _, user := range users {
		fmt.Printf("| %s | %s | %d | %s |\n", user.GetUsername(), user.GetPassword(), user.GetAge(), user.GetSex())
	}
	
	View = "indexView"
}
