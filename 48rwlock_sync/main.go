package main

import (
	"fmt"
	"sync"
	"time"
)

//rwlock 读写互斥锁

var(
	x=0
	wg sync.WaitGroup
	lock sync.Mutex
	rwlock sync.RWMutex
	sa sync.Once
)

func read(){
	defer wg.Done()
	lock.Lock()    		//加互斥锁
	//rwlock.RLock()	//读锁
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	// rwlock.RUnlock() //读解锁
	lock.Unlock()		//解互斥锁
}

func write(){
	defer wg.Done()
	//lock.Lock()	//加互斥锁
	rwlock.Lock()//写锁
	x = x+1
	time.Sleep(100*time.Millisecond)
	rwlock.Unlock()//写解锁
	//lock.Unlock()	//加互斥锁

}


func main(){
	start :=time.Now()
	for i:=0;i<10 ;i++  {
		wg.Add(1)
		go write()
	}
	//读的太快，写到1就读完，读取的时候read（）不用去等
	time.Sleep(time.Second)
	for i:=0;i<1000 ;i++  {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	end :=time.Now()
	fmt.Println(end.Sub(start))

}
