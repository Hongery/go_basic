package main

import (
	"sync"
	"fmt"
	"time"
)

//为什么需要context

var wg sync.WaitGroup
var notify bool

func f(){
	defer wg.Done()
	for {
		fmt.Println("hg")
		time.Sleep(time.Millisecond*500)
		if notify{
			break
		}
	}
}


func main(){
	wg.Add(1)
	go f()
	time.Sleep(time.Second*5)
	notify=true
	wg.Wait()

}
