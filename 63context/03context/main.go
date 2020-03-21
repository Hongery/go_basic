package main

import (
	"sync"
	"fmt"
	"time"
	"context"
)

//为什么需要context

var wg sync.WaitGroup


// 参数是context文件下的interface  Context
func f(ctx context.Context){
	defer wg.Done()
	go f2(ctx)
FORLOOP:
	for {
		fmt.Println("hg")
		time.Sleep(time.Millisecond*500)
		select {
		/*	func (*emptyCtx) Done() <-chan struct{} {return nil}*/
		case <-ctx.Done():
			break FORLOOP
		default:

		}
	}
}
func f2(ctx context.Context){
	defer wg.Done()
FORLOOP:
	for {
		fmt.Println("hg")
		time.Sleep(time.Millisecond*500)
		select {
		/*	func (*emptyCtx) Done() <-chan struct{} {return nil}*/
		case <-ctx.Done():
			break FORLOOP
		default:

		}
	}
}


func main(){
	//func WithCancel(parent Context) (ctx Context, cancel CancelFunc) cacel是空函数
	//func Background() Context  context是一个接口
	ctx,cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go f(ctx)
	time.Sleep(time.Second*5)
	//如何通知子goroutine退出
	cancel()//无论执行多少个goroutine，只要执行这个函数都会结束
	wg.Wait()

}
