package controller

import (
	"fmt"
)

type IndexController struct {

}

func (this *IndexController) Index() {
	View = "indexView"
	fmt.Println("[IndexController@Index] 暂无内容")
}

func (this *IndexController) Welcome() {
	View = "loginView"
	fmt.Println("[IndexController@Welcome]")
}