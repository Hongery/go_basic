package main

import (
	"fmt"
)

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 0
	defer calc("BB", x, calc("B", x, y))
	y = 1
	/*
	A 1 2 3
	B 0 2 2
	BB 0 2 2
	AA 1 3 4
	*/
}
//1 x=1 y=2     defer 是压栈，先进后出
//2 defer calc("AA",x,calc("A", x, y))  压栈的时候将x存进去
//3 calc("A", x, y) //"A" 1 2 3
//4 defer calc("AA",x,3)  （x =1）
//5 x=0
//6 defer calc("BB",x,calc("B", x, y))
//7 calc("B", x, y) //"B" 0 2 2    
//8 defer calc("BB",x,2)
//9 y=1 （有点混淆，没有被引用过）
//10 calc("BB",x,2) //"BB" 0 2 2  出栈  
//calc("AA",x,3) //"AA" 1 3 4   这里有点特别  （x=1 为原先的）