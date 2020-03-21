package main

import (
	"fmt"
)
//整形

func main(){
	//十进制
	var i1=101
	fmt.Printf("%d\n",i1)	//101
	fmt.Printf("%b\n",i1)  //十转二  1100101
	fmt.Printf("%o\n",i1)  //十转八 145
	fmt.Printf("%x\n",i1)	//十转十六 65
	//八进制
	i2 := 0777 
	fmt.Printf("%d\n",i2) //八转十 511
	//十六进制
	i3 := 0x1234567
	fmt.Printf("%d\n",i3) //19088743

	//查看变量的类型
	i4 :=int8(9)
	fmt.Printf("%T\n",i4) //int8
	
}