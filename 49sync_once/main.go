package main

import (
	"fmt"
	"sync"
)

//sync.Once 只执行一次操作

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
	f := func() {
		close(ch2)
	}
	once.Do(f)  //传匿名函数
	//once.Do(func() {close(ch2)}) //确保某个操作只执行一次
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
