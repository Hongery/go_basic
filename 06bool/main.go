package main

import (
	"fmt"
)

func main(){
	b1 :=true
	var b2 bool //默认是false
	fmt.Printf("%T\n",b1)
	fmt.Printf("%T value:%v\n",b1,b2)
}