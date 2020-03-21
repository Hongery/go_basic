package main

import (
	"studygo/day01/66logAgent/kafka"
	"fmt"
	"studygo/day01/66logAgent/taillog"
	"time"
	"studygo/day01/66logAgent/conf"
	"gopkg.in/ini.v1"
)

//logAgent 入口函数

var (
	cfg =new(conf.AppConf)
)
func run(){
	//1. 读取日志
for{
	select {
	case line := <- taillog.ReadChan():
	//2.发送到kafka
		kafka.SendToKafka(cfg.KafkaConf.Topic,line.Text)
	default:
		time.Sleep(time.Second)
	}
}
}

func main() {
	//0.加载配置文件
	err:=ini.MapTo(cfg,"./conf/config.ini")
	if err !=nil{
		fmt.Printf("load ini failed ,err :%v",err)
		return
	}
	//1. 初始化kafka连接
	err =kafka.Init([]string{cfg.KafkaConf.Address})
	if err !=nil{
		fmt.Println("init kafka failed",err)
		return
	}
	fmt.Println("kafka success")
	//2. 打开日志文件准备收集日志
	err =taillog.Init(cfg.TaillogConf.FileName)
	if err !=nil{
		fmt.Printf("Init taillog failed err:%v\n:",err)
		return
	}
	fmt.Println("taillog success")
	run()
}

