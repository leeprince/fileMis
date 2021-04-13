package utils

import "fmt"

type config struct {
	Key string `json: "Key"`
	Value string `json: "Value"`
}

type Config struct {
	BasePath string `json:"base_path"`
	DatePath string `json:"date_path"`
}



// TODO: [flag.String(...) 替换更优] - prince_add_todo
const configPath = "./conf/config.json"

func SetConfigData()  {
	config := &Config{}
	_ = ReadJsonFile(configPath, config)
	fmt.Printf("配置的类型：%T; 配置的值：%v", config.BasePath, config.DatePath)
}

// func GetModelData() map[string]string {
//
// }
//
// func GetModelDataPath() string {
//
// }
