package main 

import (
	"fmt"
)

//匿名函数一般在函数内部使用
func main(){
	//函数内部一般没有办法声明带名字的函数
	f1 := func(x,y int) {
		fmt.Println(x+y)
	}
	f1(10,20)

	//如果值调用一直，还可以简写成立即执行
	func (x,y int){
		fmt.Println(x+y)
		fmt.Println("hello world")
	}(100,200)
	

}