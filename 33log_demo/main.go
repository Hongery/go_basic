package main

import (
	"os"
	"fmt"
	"log"
	"time"
)

func main(){
	file , err :=os.OpenFile("./xx.log",os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err != nil{
		fmt.Printf("open file failed,err%v:\n",err)
		return
	}
	log.SetOutput(file)//写入文件中
	for {
		log.Println("这是一条测试日志")
		time.Sleep(time.Second*3)
	}
}
