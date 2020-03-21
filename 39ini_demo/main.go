package main

import (
	"fmt"
	"reflect"
	"github.com/pkg/errors"
	"io/ioutil"
	"strings"
	"strconv"
)

//ini配置文件解析器

//mysqlConfig MYSQL配置结构体
type MysqlConfig struct {
	Address 	string 	`ini:"address"`
	Port 		int		`ini:"port"`
	Username	string	`ini:"username"`
	Password 	string `ini:"password"`
}
//redisconfig
type RedisConfig struct {
	Host  string	`ini:"host"`
	Port  int		`ini:"port"`
	Password string `ini:"password"`
	Database int 	`ini:"database"`
	Test 	bool	`ini:"test"`
}
//Config
type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig	`ini:"redis"`
}
func loadIni(fileName string,data interface{}) (err error){
	//0.参数的校验
	//0.1传来的data参数必须是指针类型（因为需要在函数中对其赋值）
	t :=reflect.TypeOf(data)
	fmt.Println(t.Kind())
	if t.Kind() !=reflect.Ptr{ //ptr指针类型
		err =errors.New("data param should be a pointer")
		return
	}
	//0.2传进来的data参数必须是结构体类型指针（以为配置文件中各种键值对需要赋值给结构体字段）
	if t.Elem().Kind() !=reflect.Struct{
		err =errors.New("data param should be a struct")
		return
	}
	//1.读文件得到字节类型数据
	b,err :=ioutil.ReadFile(fileName)
	if err !=nil{
		return
	}
	//string(b) 将字节类型转换为字符串
	lineSlice :=strings.Split(string(b),"\r\n")
	//fmt.Println(lineSlice)
	//2.一行一行的读数据
	var structName string
	for idx,line := range lineSlice{
		//去掉字符串首尾的空格
		line =strings.TrimSpace(line)
		//如果是空格就跳过
		if  len(line) == 0{
			continue
		}
		//2.1如果是注释就跳过
		if strings.HasPrefix(line,";") || strings.HasPrefix(line,"#"){
			continue
		}
		//2.2如果是[开头的就表示是节（section）
		if strings.HasPrefix(line,"["){
			if line[0] !='[' || line[len(line)-1] != ']'{
				err =fmt.Errorf("line:%d syntax error",idx+1)
				return
			}
			//把这一行首尾的[]去掉，取到中间的内容把首尾的空格去掉拿到内容
			sectionName :=strings.TrimSpace(line[1:len(line)-1])
			if len(sectionName)==0{
				err =fmt.Errorf("line:%d syntax error",idx+1)
				return
			}
			//根据sectionName去掉data里面根据反射找到对应的结构体
			for i:=0;i<t.Elem().NumField();i++{
				filed :=t.Elem().Field(i)
				if sectionName == filed.Tag.Get("ini"){
					//说明找到了对应的嵌套结构体，吧字段名记下来
					structName =filed.Name
					fmt.Printf("找到%s对应的潜逃结构体%s\n",sectionName,structName)//mysql  MysqlConfig
				}
			}
		}else{
		//2.3如果不是[开头就是=分割键值对
		//1.以等号分割这一行，等号左边是key，等号右边是value
		if strings.Index(line,"=") == -1 || strings.HasPrefix(line,"="){
			err =fmt.Errorf("line:%d syntax errot",idx+1)
			return
		}
		index :=strings.Index(line,"=")
		key :=strings.TrimSpace(line[:index])
		value :=strings.TrimSpace(line[index+1:]) //自动去除空格
		//2.根据structName去data里面把对应的嵌套结构体给取出来
		v:=reflect.ValueOf(data)
		sValue :=v.Elem().FieldByName(structName)
		sType :=sValue.Type()
		if sValue.Kind() != reflect.Struct{
			err =fmt.Errorf("data中的%s字段应该是一个结构体",structName)
			return
		}
		var fieldName string
		var fileType reflect.StructField
		//3.遍历嵌套结构体的每一个字段，判断tag是不是等于key
		for i:=0 ;i<sValue.NumField();i++{
			filed :=sType.Field(i)
			fileType =filed
			if filed.Tag.Get("ini") ==key{
				//找到相应的字段
				fieldName= filed.Name
				break
			}
		}
		//4.如果key=tag ，给这个字段赋值
		//4.1根据fieldName去取出这个字段
		if len(fieldName) == 0 { //value长度为0
			//在结构体中找不到对应的字符
			continue
		}
		fileObj :=sValue.FieldByName(fieldName)
		//4.2对其赋值
		fmt.Println(fieldName,fileType.Type.Kind())//address string
			switch fileType.Type.Kind() {
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int,reflect.Int8,reflect.Int16,reflect.Int32,reflect.Int64:
				var valueInt int64
				valueInt,err = strconv.ParseInt(value,10,64)//10进制，64位
				if err != nil {
					err =fmt.Errorf("ling:%d value type error ",idx+1)
					return
				}
				fileObj.SetInt(valueInt)
			case reflect.Bool:
				var valueBool bool
				valueBool,err = strconv.ParseBool(value)
				if err !=nil {
					err =fmt.Errorf("line:%d value type error",idx+1)
					return
				}
				fileObj.SetBool(valueBool)
			case reflect.Float32,reflect.Float64:
				var valueFloat float64
				valueFloat,err = strconv.ParseFloat(value,64)
				if err != nil {
					err =fmt.Errorf("line:%d value type error",idx+1)
					return
				}
				fileObj.SetFloat(valueFloat)
			}
		}
	}
	return
}

func main(){
	var cfg Config
	err :=loadIni("./conf.ini",&cfg)
	if err !=nil {
		fmt.Printf("load ini failed,err %v\n",err)
		return
	}
	fmt.Printf("%#v\n",cfg)
}