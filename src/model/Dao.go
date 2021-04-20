package model

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// 统一模型接口
type Model interface {
	ToString() string
}

// 读取表（文件）数据
func RfData(table string, key string, data map[string]Model) error {
	// 读取"表"文件，包含文件路径及文件名
	fileName := path + table + subfix
	f, ferr := os.Open(fileName)
	if ferr != nil {
		panic("读取文件失败。file:" + fileName)
	}

	// 延迟执行：关闭文件资源流
	defer f.Close()

	// 创建读取的文件的缓冲区
	buf := bufio.NewReader(f)

	/* 遍历文件数据
	分隔符
		列：每一行以英文逗号分隔
		字段：第一行为字段
		数据：非第一行，以换行符为分隔
	*/
	field := make([]string, 0)
	rowNum := 0
	for {
		row, rerr := buf.ReadBytes('\n')
		rowNum++
		fmt.Println("读取第几行：", rowNum)
		if rerr != nil {
			if rerr == io.EOF { // 结束行。原文件最后一行数据数据之后应换行，否则读取不到最后一行的数据
				fmt.Println("读取文件结束符")
				break
			}
			panic("读取文件-行失败。file:" + fileName)
		}
		rowData := strings.Split(strings.TrimSuffix(string(row), "\n"), ",")
		if len(field) == 0 { // 第一行存储字段
			field = rowData
			fmt.Println("读取到的字段为：", field)
		} else { // 非第一行存储数据
			fmt.Println("读取非字段行的行数据为：", rowData)
			merr := toModel(table, key, data, rowData, field)
			return merr
		}
	}

	return nil
}

/* 存储数据到 models
1. 根据 table 得到模型
2. 利用反射对模型进行赋值
3. 再把模型存储到data中
*/
func toModel(table string, key string, data map[string]Model, rowData []string, field []string) error {
	// 1. 根据 table 得到模型
	if _, ok := sysModels[table]; !ok {
		return errors.New("需要存储的模型不存在")
	}
	fmt.Printf("得到的模型。原数据：%v, 类型：%T \n", sysModels[table], sysModels[table])
	
	// 2. 利用反射对模型进行赋值
	// 获取模型的结构体
	arg := []reflect.Value{}
	modelValue := reflect.ValueOf(sysModels[table])
	modelStructFunc := modelValue.Call(arg) // Call方法使用输入的参数in调用v持有的函数:func (v Value) Call(in []Value) []Value
	fmt.Printf("Reflect Call 后的类型：%T ;原值：%v, \n", modelStructFunc, modelStructFunc)
	modelFunc := modelStructFunc[0]
	fmt.Printf("Reflect Call 取[0]的类型：%T ;原值：%v, \n", modelFunc, modelFunc)
	
	var keyValue string
	for i, colDat := range rowData {
		// 查询字段的值
		if field[i] == key {
			keyValue = colDat
		}
		// 获取模型中对应设置相应字段的方法
		methodName := "Set" + strings.Title(field[i])
		fmt.Println("拼接成的方法名:"+methodName)
		
		fieldSetFunc := modelFunc.MethodByName(methodName)
		fieldSetFunc.Call([]reflect.Value{
			reflect.ValueOf(toTypeValue(modelFunc, field[i], colDat)),
		})
		
		/*getMethodName := "Get" + strings.Title(field[i])
		fieldGetFunc := modelFunc.MethodByName(getMethodName)
		getValue := fieldGetFunc.Call([]reflect.Value{})
		fmt.Printf("%s 之后 %s 的 类型：%T;值：%v \n", methodName, getMethodName, getValue[0], getValue[0])*/
	}
	fmt.Printf("查询的主键：%v;主键值：%v \n", key, keyValue)
	fmt.Printf("modelFunc.Interface() 类型：%T;值：%v \n", modelFunc.Interface(), modelFunc.Interface())
	
	data[keyValue] = modelFunc.Interface().(Model)

	return nil
}

func toTypeValue(modelV reflect.Value, field, value string) interface{} {
	mtype := modelV.Elem().FieldByName(field).Type().Name() // [<推荐] // 另外一种写法：mtype := reflect.Zero(modelV.Type().Elem()).FieldByName(field).Type().Name()
	switch mtype {
	case "int":
		b, _ := strconv.Atoi(value)
		return b
	}
	return string(value)
}
