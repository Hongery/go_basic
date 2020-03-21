package main

import (
	"fmt"
)


func main(){
	//元素类型为map切片
	var s1=make([]map[int]string,10,20)//key=int value=string
	//没有对内部的map做初始化
	s1[0]=make(map[int]string,1) //但是s1[0] 为一个map  有10个map
	s1[0][10]="shazi"
	fmt.Println(s1)//[map[10:shazi] map[] map[] map[] map[] map[] map[] map[] map[] map[]]

	//值为类型切片的map ,要记得初始化
	var m1 = make(map[string][]int, 10)  //key=string  value=[]int
	m1["北京"] =[]int{10,20,30} //map[北京:[10 20 30]]
	fmt.Println(m1)


	//元素类型为map的切片
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
		fmt.Println(index,value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "小王子"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "沙河"
	for _, value := range mapSlice {
		fmt.Println( value)
	}

	m :=make(map[string]int,8)
	m["user"] =1
	m["is"]=2
	m["s"]=2
	for key,value :=range m{
		fmt.Println(key,value)
	}

}