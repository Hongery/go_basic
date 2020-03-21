package main

import (
	"fmt"
)

//同一个结构体可以是新多个接口
//接口还可以嵌套
type animal interface {
	mover
	eater
}
type mover interface{
	move()
}
type eater interface{
	eat(string)
}

type cat struct {
	name string
	feet int8
}

//使用指针接收
func (c *cat) move() {
	fmt.Println("miao")
}
func (c *cat) eat(food string) {
	fmt.Printf("猫吃%s\n", food)
}
func main() {
	var a1 animal
	c1 := cat{"tom", 4}   //{"tom",4}
	c2 := &cat{"jack", 5} //&{"jack",5}

	 //a1 = c1 使用指针会报错 ，实现animal这个接口是cat的指针类型
	a1 =&c1
	 fmt.Println(a1)
	a1 = c2
	fmt.Println(a1)
}