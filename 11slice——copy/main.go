package main

import (
	"sort"
	"fmt"
)

func main(){
	a1 :=[]int{1,3,5}
	a2 := a1  //赋值
	// var a3 []int //这样定义a3为nil  没办法进行copy
	var a3 =make([]int, 3,3)
	copy(a3,a1)  //a1复制到a3
	fmt.Println(a1,a2,a3)//[1 3 5] [1 3 5] [1 3 5]
	a1[0]=100
	fmt.Println(a1,a2,a3)//[100 3 5] [100 3 5] [1 3 5]  由于切片是引用类型a1给a2赋值，都指向同一块内存地址

	x1 :=[...]int{1,2,3} //数组
	s1 :=x1[:]		//切片
	fmt.Println(s1,len(s1),cap(s1))//[1 2 3] 3 3 
	s1=append(s1[:1],s1[2:]...)
	fmt.Println(s1,len(s1),cap(s1))//[1 3] 2 3 删除一个元素
	fmt.Println(x1) //[1 3 3]  //修改了底层数组

	var a =make([]int,5,10)
	fmt.Println(a)//[0 0 0 0 0]
	for i :=0;i<10;i++{
		a =append(a,i)
	}
	fmt.Println(a)//[0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]
	fmt.Println(cap(a)) //20扩容 面试题

	var b1 =[...]int{4,2,1,9}
	sort.Ints(b1[:]) //对切片进行排序
	fmt.Println(b1)//1，2，4，9





}