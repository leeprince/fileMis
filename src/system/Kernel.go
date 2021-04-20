package system

import (
	"fileMis/src/utils"
	"fmt"
)

func Run()  {
	for {
		flag := utils.CReturn(dispatch)
		if flag {
			break
		}
	}

	fmt.Println("结束")
}
