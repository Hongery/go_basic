package main

import (
	"fmt"
)
//常量，定义之后不会修改，程序在运行期间不会改变的量
var pi =3.1415926
//批量声明常量
const(
	statuOk=200
	notFount=400
)
//const同时声明多个常量时，如果省略了值则表示和上面一行的值相同
const(
	n1=100
	n2
	n3
)
func main(){
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
}