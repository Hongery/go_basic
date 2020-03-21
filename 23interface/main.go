package main

import (
	"fmt"
)

//引出一个能叫的类型
type speaker interface{
	speak()  //只要实现speak方法的变量都是speaker类型
}

type cat struct{}
type dog struct{}
type person struct{}

func (c cat) speak(){
	fmt.Println("miao ")
}
func (c dog) speak(){
	fmt.Println("wang ")
}
func (c person) speak(){
	fmt.Println("a  ")
}
func do(c speaker){
	c.speak()
}

func main(){
	var a cat 
	var b dog
	var c person
	do(a)
	do(b)
	do(c)
	

}