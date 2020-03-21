package main

import (
	"sync"
	"fmt"
	"time"
)

//为什么需要context

var wg sync.WaitGroup
var exitChain=make(chan bool,1)

func f(){
	defer wg.Done()
FORLOOP:
	for {
		fmt.Println("hg")
		time.Sleep(time.Millisecond*500)
		select {
		case <-exitChain:
			break FORLOOP
		default:

		}
	}
}


func main(){
	wg.Add(1)
	go f()
	time.Sleep(time.Second*5)
	//如何通知子goroutine退出
	exitChain <-true
	wg.Wait()

}
