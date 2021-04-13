package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

/*{
    "ModelData": {
        "Path": "./src/data",
        "Subfix": ".sql"
    }
}*/
func ReadJsonFile(path string, data interface{}) error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("读取Json文件错误")
		return err
	}
	fmt.Printf("读取Json文件的[]byte数据：%v \n", bytes)
	jerr := json.Unmarshal(bytes, data)
	if jerr != nil {
		fmt.Println("解析Json文件数据错误")
		return jerr
	}
	fmt.Printf("解析Json文件的数据：%v \n", data)
	return nil
}
