# go 基于文件存储的用户管理系统

## 配置文件说明
1. 配置文件路径：conf/config.json
2. 用户信息的文件路径：ModelData.Path; 文件名后缀：ModelData.Subfix

> 该路径文件不存在时会自动创建初始化数据

> 默认的用户信息：
```
username,password,sex,age
prince,prince,男,18

```

> 初始化的用户信息：
```
username,password,sex,age
prince,123456,男,18

```

## 运行
项目根目录下执行：go run src/main.go