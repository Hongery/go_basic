package main

import (
	"fmt"
)

func assign(x interface{}){
	fmt.Printf("%T\n",x)
	v, ok := x.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}
}

func assign2(x interface{}){
	fmt.Printf("%T\n",x)
	switch t :=x.(type) {
	case string:
		fmt.Println("是一个字符串",t)
	case int:
		fmt.Println("是一个int",t)
	case int64:
		fmt.Println("是一个int64",t)
	case bool:
		fmt.Println("是一个bool",t)	
	
	}	
	
}

func main(){
	assign(100) //报错 interface int ，not string
	assign2(100)
}