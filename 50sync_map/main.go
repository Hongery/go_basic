package main

import (
	"sync"
	"strconv"
	"fmt"
)
//map 并发不是安全的
//goroutine map 安全问题  会出现
//fatal error: concurrent map writes
//添加完互斥锁也会出现问题
//fatal error: concurrent map read and map write
//使用sync.Map是可以的

//var m  = make(map[string]int)
//var lock sync.Mutex
/*func get(key string) int {
	return m[key]
}
func set(key string, value int) {
	m[key] = value
}*/
/*
func main() {
	wg :=sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)//数字转换为字符串
			lock.Lock()
			set(key, n)
			lock.Unlock()
			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}
*/
//sync内置map类型
var m = sync.Map{}
func main(){
	wg :=sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)//数字转换为字符串
			m.Store(key,n)//必须使用sync.Map内置的Store方法设置键值对
			value,_:=m.Load(key)//必须使用sync.Map提供的load方法根据key取值
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
