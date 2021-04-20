package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CInfo interface {
	call() (bool, error)
}
type Cfunc func() (bool, error)

func (c Cfunc) call() (bool, error) {
	return c()
}

var	(
	inputReader *bufio.Reader
)

const ESC_CHAR string = "x"

func init() {
	inputReader = bufio.NewReader(os.Stdin)
}

// 格式化输出
func CReturn(a Cfunc) bool {
	fmt.Println("⬇︎⬇︎⬇︎⬇︎⬇︎⬇︎⬇︎⬇︎⬇︎⬇︎⬇︎⬇︎⬇︎⬇︎⬇︎⬇︎⬇︎⬇︎⬇︎⬇︎⬇︎⬇︎⬇︎⬇︎⬇︎⬇︎⬇︎⬇︎⬇︎⬇︎⬇︎")
	flag, err := a.call() // == flag, err := a()
	if err != nil {
		fmt.Println("错误信息：", err)
	}
	fmt.Println("⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎⬆︎")
	fmt.Println("")
	return flag
}

// 读取用户输入的内容
func CRead() string {
	input, _ := inputReader.ReadString('\n')
	input = strings.TrimSpace(strings.TrimSuffix(input, "\n"))
	return input
}

// 输出下一步操作的信息
func CReturnNextOperDesc(d []string)  {
	for i, i2 := range d {
		fmt.Printf("(%d) : %s \n", i, i2)
	}
	fmt.Printf("> 退出请键入：control+c 或者输入：%s \n", ESC_CHAR)
}