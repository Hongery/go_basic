package main

import (
	"fmt"
	"bufio"
	"os"
)

//获取用户输入时如果有空格
func useScan(){
	var s string
	fmt.Print("请输入内容")
	fmt.Scanln(&s)  // a b c   只能有a 空格结束读取
	fmt.Printf("你输入的内容是：%s \n",s)
}

//换行结束
func useBufio(){
	var s string
	fmt.Print("请输入内容")
	reader := bufio.NewReader(os.Stdin) //表示输入，并获取对象
	s,_=reader.ReadString('\n') //换行符结束
	fmt.Printf("你输入的内容是：%s\n",s)
}

func main(){
	//useScan()
	useBufio()
}
