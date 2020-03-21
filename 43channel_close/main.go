package main

import (
	"fmt"
)


//关闭通道

//close(ch1) 与range的使用
func main(){
	ch1 :=make(chan int,2)
	ch1 <- 10
	ch1 <- 20
	close(ch1)
	// for x :=range ch1{
	// 	fmt.Println(x)
	// }
	<-ch1
	<-ch1
	x,ok := <-ch1  //已经取出两个值
	//对已经关闭的通道去取值是可以取到的
	fmt.Println(x,ok)  // 0 false
	
}
