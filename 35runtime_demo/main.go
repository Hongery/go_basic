package main

import (
	"runtime"
	"fmt"
	"path"
)

func f(){
	pc,file,line,ok :=runtime.Caller(1)//数字表示被调用的层数,向上
	if !ok{
		fmt.Printf("runtime.Caller() failed \n")
		return
	}
	funcName :=runtime.FuncForPC(pc).Name()
	fmt.Println(funcName) //函数方法名 main.f1
	fmt.Println(file) //文件名 c:/Program Files/StudyGo/src/studygo/day01/35runtime_demo/main.go
	fmt.Println(path.Base(file)) //获取最后的文件 main.go
	fmt.Println(line)//行号 21
}

func f1(){
	f() //21
}

func main(){
	f1()
}
