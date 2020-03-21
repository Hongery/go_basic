package main

import (
	"studygo/day01/mylogger"
	"time"
)
var log mylogger.Logger

//测试我们自己写的日志库
func main() {
	log = mylogger.NewConsoleLogger("INFO") //终端日志
	log = mylogger.NewFileLogger("INFO", "./", "hello.log", 10*1024)//写入文件日志
	for {
		log.Debug("这是一条Debug日志")
		log.Info("这是一条Info日志")
		log.Warning("这是一条Warning日志")
		id := 10010
		name := "lixiang"
		log.Error("这是一条Error日志id:%d,name:%s", id, name)
		log.Fatal("这是一条Fatal日志")
		time.Sleep(time.Second * 3)
	}
}
