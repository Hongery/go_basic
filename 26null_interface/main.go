package main

import (
	"fmt"
)

//空接口

//interface: 关键字
//interface{}:空接口类型

//空接口做函数的参数
func show(a interface{}) {
	fmt.Printf("type:%T  value:%v\n",a,a)
}

func main() {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	//空接口可以接受任何类型
	m1["name"] = "周琳"
	m1["age"] = 900
	fmt.Println(m1)

	show(false)
	show(nil)
	show(m1)
}
