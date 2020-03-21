package main

import (
	"strconv"
	"fmt"
)

func main(){
	//从字符串中解析出对应的数据
	str :="10000"
	ret1,err :=strconv.ParseInt(str,10,64) //10进制  64位
	if err !=nil {
		fmt.Println("parseint failed err:",err)
		return
	}
	fmt.Printf("%#v %T\n",ret1,ret1) //10000   int64 字符串转换成数字

	strconv.Atoi()
	//把数字转换成字符串类型
	i := int32(97)
	ret2 :=fmt.Sprintf("%d",i)
	fmt.Printf("%#v",ret2)  //"97"
}
