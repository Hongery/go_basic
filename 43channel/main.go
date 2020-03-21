package main

import (
	"fmt"
	"sync"
)

//channel 通道必须初始化才能使用

var a []int
//chan 是引用类型默认为nil，需要初始化才能使用，make/new开辟内存空间
var b chan int //需要指定通道中元素的类型
var wg sync.WaitGroup

//不带缓冲区通道
func noBufChannel(){
	fmt.Println(b) //nil
	b =make(chan int) //不带缓冲区通道的初始化
	wg.Add(1)
	//b <-10  //all goroutines are asleep - deadlock!
	go func() {
		defer wg.Done()
		x :=<-b
		fmt.Println("后台goroutine从通道b中取",x)
	}()
	b <- 10
	fmt.Println("10 发送到通道b中了")
	wg.Wait()
}

//带缓冲区通道
func bufChannel(){
	fmt.Println(b)
	b =make(chan int, 16) //带缓冲区通道的初始化
	b <- 10 //发送
	fmt.Println(b)//打印地址
	fmt.Println("10发送到b通道")
	x :=<-b //接收
	fmt.Println("从通道中取到了",x)
}

func main(){
	// noBufChannel()
	bufChannel()
}
