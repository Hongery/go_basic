package main

import (
	"encoding/json"
	"fmt"
)

//1序列化 把go语言结构体变量————转换成json格式的字符串
//2.反序列化  将json的字符串--go语言中能识别的结构体变量

type person struct {
	Name string `json:"name" db:"name" ini:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "周琳",
		Age:  9000,
	}
	//序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("marshal failed")
		return
	}
	fmt.Printf("%v\n", string(b))

	//反序列化
	str := `{"name":"liiang","age":18}`
	var p2 person
	json.Unmarshal([]byte(str), &p2) //传指针是为了能在json。Unmarshal内部修改person
	fmt.Printf("%#v\n", p2)

}
