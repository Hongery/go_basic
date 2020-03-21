package main

import (
	"fmt"
)

//声明变量
var name string

//批量声明
var (
	n    string
	age  int
	isOk bool
)

func main() {
	name = "lixiang"
	age = 16
	n = "a"
	isOk = true
	//GO语言声明变量必须使用，不使用编译不过去
	fmt.Print(isOk)             //在终端中输出
	fmt.Printf("name:%s", name) //%s:占位符 使用name这个变量去替换占位符
	fmt.Println(age)            //打印完后会在后面加上一个换行符
	//声明变量同时赋值
	var s1 string = "hg"
	fmt.Println(s1)
	//类型推导
	var s2 = "20"
	fmt.Println(s2)
	//简短变量声明，声明并初始化，只能在函数内部使用
	s3 := "hahah"
	fmt.Println(s3)
	//s1 :=10 同一个作用域{} 中不能重复生命同样的变量

}
