package model

import (
	"fileMis/src/utils"
	"fmt"
)

var (
	needInitDataByTableSlice []string = []string{userModelTable}
	path                     string   = utils.GetConfig().GetModelDatePath()   // 表的存储路径。当前路径为：执行go命令时的当前路径
	subfix                   string   = utils.GetConfig().GetModelDateSubfix() // 相应表的后缀
	sysModels                map[string]interface{}
)

var (
	needInitData map[string]string = map[string]string{
		userModelTable: "username,password,sex,age\nprince,123456,男,18\n",
	}
)

func init() {
	initData()
	initModel()
}

// 初始化模型
func initModel() {
	sysModels = make(map[string]interface{})
	sysModels[userModelTable] = NewUserModel
	userData = make(map[string]Model, 0)
	err := RfData(userModelTable, userDataKey, userData)
	if err != nil {
		fmt.Println("初始化模型的表数据失败。err:", err)
	}
}

// 初始化数据
func initData() {
	for _, table := range needInitDataByTableSlice {
		checkInitDataTableAndCreate(table)
	}
}

// 检查需要初始化的文件是否存在，不存在则重新写入
func checkInitDataTableAndCreate(table string) {
	tablePath := path + table + subfix
	bool, err := utils.PathExist(tablePath)
	if !bool {
		fmt.Println("初始化", tablePath, "的数据信息失败。err:", err)
		// 创建文件并写入初始化数据
		tableData, bool := needInitData[table]
		if ! bool {
			fmt.Println("需要初始化", tablePath, "的数据信息不存在")
		}
		bool, err := utils.WrtiteFile(path, table + subfix,  tableData)
		if ! bool || err != nil {
			fmt.Println("写入初始化数据失败", ";boll:", bool, ";err:", err)
		}
		fmt.Printf("重新写入初始化数据成功。路径：%s data: \n%s \n", tablePath, tableData)
	}
}
