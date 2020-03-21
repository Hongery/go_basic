package main

import (
	"fmt"
)

//命名的函数值就相当于在函数中声明一个变量
func f(x,y int)(res int){  //若func f(x,y int) int(){ return x+y}
	res =x+y
	return //命名的返回值可以不写res
} 

//可变长参数，必须放在函数的最后
func s(s string,y... int){
	fmt.Println(s)
	fmt.Println(y)  //y的类型是切片slice
}
//go函数没有默认参数这个概念

func main(){
	r :=f(3,4)
	fmt.Println(r)

	s("hg")   //hg []
	s("huang",1,2,3,4)//huang [1 2 3 4]
}