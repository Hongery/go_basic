package main

import (
	"math/rand"
	"fmt"
	"time"
	"sync"
)

func f(){
	rand.Seed(time.Now().UnixNano())//随机种子，保证每次执行的时候都有点不一样
	for i:=0;i<5;i++{
		r1 :=rand.Int()
		r2 :=rand.Intn(10)//0<=x<10 [0,10)
		fmt.Println(r1,r2)
	}
}

func f1(i int){
	defer wg.Done()//计数器减一
	time.Sleep(time.Millisecond*time.Duration(rand.Intn(300)))
	fmt.Println(i)

	}

var wg sync.WaitGroup //定义全局变量

func main(){
	//wg.Add(10)下面的方式也可以
	for i:=0 ;i<10;i++{
		wg.Add(1)
		go f1(i)
	}
	//如何知道这10个routine都结束了
	wg.Wait() //等待计数器减为0
}

