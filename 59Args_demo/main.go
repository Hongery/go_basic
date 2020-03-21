package main

import (
	"fmt"
	"os"
)

//os.Args 获取命令行参数   ./59Args_demo.exe a b c(这三个是参数)
func main() {
	fmt.Printf("%#v\n", os.Args)        //[]string{"59Args_demo.exe","a","b","c"}
	fmt.Println(os.Args[0], os.Args[2]) // a ,c
	fmt.Printf("%T\n", os.Args)         //类型 []string
}
