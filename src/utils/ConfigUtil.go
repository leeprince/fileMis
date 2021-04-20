package utils

import (
	"flag"
	"fmt"
)

var (
	instance *uconfig
	// 查看帮助命令：go run src/main.go -h
	conf     = flag.String("conf", "./conf/config.json", "fileMis`s config.json customer path")
)

// 结构体中的字段除了有名字和类型外，还可以有一个可选的标签（tag）：它是一个附属于字段的字符串，可以是文档或其他的重要标记
// json 解析时只有字段为可访问的才能转换成功
type uconfig struct {
	BasePath string `json:"BaseRoot"` // 结构体的字段标签定义的json键必须等于 json 结构中定义的键；读取结构体时仍然为结构体字段
	ModelData ModelData
}

type ModelData struct {
	Path string // json的键默认是结构体字段
	Subfix string // json的键默认是结构体字段
}

func init() {
	flag.Parse()
	NewUConfigWithFile(*conf)
}

func NewUConfigWithFile(name string) *uconfig {
	if instance == nil {
		c := &uconfig{}
		ReadJsonFile(name, c)
		instance = c // <--- NOT THREAD SAFE
		fmt.Println("读取到的配置：", c.BasePath, c.ModelData, c.ModelData.Path, c.ModelData.Subfix)
	}
	
	return instance
}

func GetConfig() *uconfig {
	return instance
}
func (u *uconfig) GetModelDatePath() string {
	fmt.Println("读取到的配置路径为：", u.ModelData.Path)
	return u.ModelData.Path
}
func (u *uconfig) GetModelDateSubfix() string {
	fmt.Println("读取到的文件后缀为：", u.ModelData.Subfix)
	return u.ModelData.Subfix
}
