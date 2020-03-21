package main

import (
	"fmt"
)


func modifyArray(x *[3]int) {
	x[0] = 100
}
func modifyArary3(x [3][2]int){
	x[2][0]=100
}
func modifyArray2(x [][]int) {
	x[2][0] = 100
}
func main() {
	a := [3]int{10, 20, 30}
	modifyArray(&a) //在modify中修改的是a的副本x
	fmt.Println(a) //[10 20 30]
	b := [][]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	modifyArray2(b) //在modify中修改的是b的副本x在modifyArray3中
	fmt.Println(b)  //[[1 1] [1 1] [1 1]]
	fmt.Println(b)  ////[[1 1] [1 1] [100 1]]  modifyArray2中  var a []int 不定长数组成为切片
}