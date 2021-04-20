package system

import (
	"errors"
	"fileMis/src/controller"
	"fileMis/src/utils"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func dispatch() (bool, error) {
	// 1. 获取控制器和展示的列表
	args := strings.Split(controller.Next, "::")
	// fmt.Println("获取控制器和展示的列表的结果为：", args)
	ctl, ok := controller.Controllers[args[0]]
	if ok != true {
		fmt.Println("无法根据指定标识查找控制器", args[0])
	}
	// 2. 执行方法
	// fmt.Printf("controller.Controllers: %v ; ctl reflect.TypeOf:%s ; ctl reflect.ValueOf:%s \n", controller.Controllers, reflect.TypeOf(ctl), reflect.ValueOf(ctl))
	methodName := reflect.ValueOf(ctl).MethodByName(args[1])
	// fmt.Printf("controller.Controllers.methodName: %v ;args[1]: %s \n", reflect.TypeOf(methodName), args[1])
	methodName.Call([]reflect.Value{})
	// 3. 获取下一步需要操作的方法
	view, ok := controller.Views[controller.View]
	if !ok {
		fmt.Println()
		return false, errors.New("无法根据指定标识查找到视图 > controller.View: " + controller.View)
	}
	// 4. 处理输出格式
	fmethod, fdesc := controller.ViewsFormate(view)
	
	// 输出可以进行下一步操作的信息
	utils.CReturnNextOperDesc(fdesc)
	// 5. 匹配输入的操作指令，分发下一步执行的方法
	for {
		inCommand := utils.CRead()
		if inCommand == utils.ESC_CHAR {
			return true, nil
		}
		flag, err := strconv.Atoi(inCommand)
		if err == nil && flag < len(fmethod) {
			controller.Next = fmethod[flag]
			break
		}
		fmt.Println("输入无效指令，请重新输入")
	}
	
	return false, nil
}
