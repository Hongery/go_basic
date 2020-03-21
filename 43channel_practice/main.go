package main

import (
	"fmt"
	"sync"
)

//channel练习
//1.启动一个goroutine，生成10个数发送到ch1
//2.启动一个goroutine，从ch1中取值，计算其平方放到ch2
//3.在main中 从ch2取值打印出来
var wg sync.WaitGroup
var once sync.Once

func f1(ch1 chan int){
	defer wg.Done()
	for  i:=0;i <10;i++{
		ch1 <- i
	}
	close(ch1)
}

func f2(ch1,ch2 chan  int){
	defer wg.Done()
	for x := range ch1{
		ch2 <- x*x
	}
	//close(ch2)
	once.Do(func() {close(ch2)})
}



func main(){
	a := make(chan int,20)
	b := make(chan int,20)
	wg.Add(3)
	go f1(a)
	go f2(a,b)
	go f2(a,b)
	wg.Wait()
	for ret := range b{
		fmt.Println(ret)
	}

}
