package main

import (
	"fmt"
)

func main(){
	n :=18
	p :=&n
	fmt.Println(&n)//0xc04200e0a0
	fmt.Printf("%T\n",p) //*int
	fmt.Println(p) //0xc04200e0a0
	fmt.Println(*p)//18

	m :=*p
	fmt.Println(m) //18
	fmt.Printf("%T\n",m)//int


	// var a *int //nil pointer
	// *a =100
	// fmt.Println(*a)//报错无效内存地址和空指针

	var a1 *int  
	fmt.Println(a1)  //nil
	var a2=new(int)
	fmt.Println(a2) //0xc04204a0e0
	*a2 =100
	fmt.Println(*a2) //100
	
}