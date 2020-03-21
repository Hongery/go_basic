package main

import (
	"fmt"
)

func f(x, y int) {
	fmt.Println("this is f")
	fmt.Println(x + y)
}

func main() {
	res := lixiang(f, 100, 200)
	res()
}

func lixiang(x func(int, int), m, n int) func() {
	//1方法
	/*tmp := func() {
		x(m, n)
	}
	return tmp*/

	//2方法
	return func(){
		x(m,n)
	}
}
