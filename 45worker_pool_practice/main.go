package main

import (
	"math/rand"
	"time"
	"sync"
	"fmt"
)
/*	实现一个计算int64随机数各位数和的程序。
	1. 开启一个`goroutine`循环生成int64类型的随机数，发送到`jobChan`
	2. 开启24个`goroutine`从`jobChan`中取出随机数计算各位数的和，将结果发送到`resultChan`
	3. 主`goroutine`从`resultChan`取出结果并打印到终端输出
*/

type job struct {
	value int64
}

type result struct {
	job	*job
	sum int64
}

var jobChan =make(chan *job,100)
var resultChan =make(chan *result,100)
var wg sync.WaitGroup
func rom(z1 chan<- *job){
	wg.Done()
	//1.循环生成int64类型的随机数，发送到`jobChan`
	for  {
		 x := rand.Int63()
		newJob := &job{
			x,
		}
		z1 <- newJob
		time.Sleep(time.Millisecond*500)
	}
}
func get(z1 <-chan *job,resultChan chan<- *result){
	wg.Done()
	//2.从`jobChan`中取出随机数计算各位数的和，将结果发送到`resultChan`
	for   {
		job := <- z1
		sum := int64(0)
		n :=job.value
		for n>0 {
			sum += n%10
			n=n/10
		}
		newResult := &result{
			job,
			sum,
		}
		resultChan<-newResult
	}
}

func main(){
	wg.Add(1)
	go rom(jobChan)
	wg.Add(24)
	//开启24个goroutine执行get方法
	for i := 0;i<24;i++{
		go get(jobChan,resultChan)
	}
	//3.从`resultChan`取出结果并打印到终端输出
	for result := range resultChan{
		fmt.Printf("value:%d  sum:%d\n",result.job.value,result.sum)
	}
	wg.Wait()
}