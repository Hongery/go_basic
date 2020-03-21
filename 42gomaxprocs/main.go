package main

import (
	"sync"
	"fmt"
	"runtime"
)

//gomaxprocs 最大线程数

var wg sync.WaitGroup
func a(){
	defer  wg.Done()
	for i:=0;i<10;i++ {
		fmt.Printf("A:%d\n",i)
	}
}
func b(){
	defer wg.Done()
	for i:=0;i<10 ;i++  {
		fmt.Printf("B:%d\n",i)
	}
}

func main(){
	runtime.GOMAXPROCS(6)//设置多少个线程同时执行程序 默认CPU逻辑核心数是6，默认跑满整个cpu
	fmt.Println(runtime.NumCPU()) //最大线程数是4，我的电脑
	wg.Add(2)
	go a()
	go b()
	wg.Wait() //等待线程数为0
}

