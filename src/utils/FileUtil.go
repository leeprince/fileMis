package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadJsonFile(path string, data interface{}) error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("读取Json文件错误")
		return err
	}
	// fmt.Printf("读取Json文件的[]byte数据：%v \n", bytes)
	jerr := json.Unmarshal(bytes, data)
	if jerr != nil {
		fmt.Println("解析Json文件数据错误")
		return jerr
	}
	// fmt.Printf("解析Json文件的数据：%v \n", data)
	return nil
}

func PathExist(path string) (bool, error)  {
	if _, err := os.Stat(path); err != nil {
		return false, err
	}
	return true, nil
}

func WrtiteFile(path, file, data string) (bool, error)  {
	// 判断是否已存在目录
	if exist, _ := PathExist(path); !exist {
		// 不存在则创建
		os.Mkdir(path, os.ModePerm) // 0777
	}
	// 打开文件 4：读；2：写；1：可执行；文件不存在则创建
	outputFile, outputErr := os.OpenFile(path + file, os.O_WRONLY|os.O_CREATE, 0666)
	if outputErr != nil {
		return false, outputErr
	}
	defer outputFile.Close()
	// 创建写的缓存
	outputWriter := bufio.NewWriter(outputFile)
	// 写入信息
	outputWriter.WriteString(data)
	// 刷新文件缓存
	outputWriter.Flush()
	return true, nil
}