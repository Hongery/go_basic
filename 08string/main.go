package main

import (
	"strings"
	"fmt"
)

func main(){
	//多行字符串
	s1 := `第一行
	第二行
	第三行
	`
	fmt.Println(s1)

	//字符串拼接
	name :="hah"
	world :="shf"
	ss :=name +world
	fmt.Println(ss)

	sss:=fmt.Sprintf("%s%s",name,world)   //return string
	fmt.Println(sss)
	// fmt.Printf("%s%s",name,world)

	//分割

	ret :=strings.Split(sss,"h")  //[ hahs hf]
	fmt.Println(ret) // [ a s f]

	s4 :="abcdef"
	fmt.Println(strings.Index(s4,"a"))


}