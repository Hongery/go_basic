package main

import (
	"fmt"
	"math"
)
func main() {
	f1 :=1.2345
	fmt.Printf("%T\n",f1)//go语言默认中的小数都是float64类型
	fmt.Printf("%f\n", math.Pi)//3.141593
	
	fmt.Printf("%.2f\n", math.Pi)//3.14
}