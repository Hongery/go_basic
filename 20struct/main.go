package main

import (
	"fmt"
)

type Person struct{
	age int
}
//构造函数
func newPerson(age int) *Person{
	return &Person{
		age : age,
	}
}

//使用值接收者：传拷贝进去
func (p Person) add1(){
	p.age++
}

//指针接收者：穿内存地址
func (p *Person) add2(){
	p.age++
}
func main(){
	p1 :=newPerson(2)
	p1.add1()
	fmt.Println(p1.age)//2

	p2 :=newPerson(2)
	p2.add2()
	fmt.Println(p2.age)//3
}