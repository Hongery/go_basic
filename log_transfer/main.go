package main

import (
	"fmt"
	"studygo/day01/log_transfer/conf"
	"studygo/day01/log_transfer/es"
	"studygo/day01/log_transfer/kafka"

	"gopkg.in/ini.v1"
)

//log transfer  运行得开启68logAgent
//将日志数据从kafka取出来发往ES

func main() {
	//0.加载配置文件
	var cfg conf.LogTransfer //下面要传指针
	err := ini.MapTo(&cfg, "./conf/cfg.ini")
	if err != nil {
		fmt.Printf("ini confg failed.err :%v", err)
		return
	}
	fmt.Printf("cfg:%v\n", cfg)
	//1.初始化ES
	//1.1初始化一个ES连接的client
	//1.2对外提供一个往ES写入数据的一个函数
	err = es.Init(cfg.ESCfg.Address,cfg.ESCfg.ChanSize,cfg.ESCfg.Nums)
	if err != nil {
		fmt.Printf("init es failed err :%v\n", err)
		return
	}
	fmt.Println("es success")
	//2.初始化kafkaf
	//2.1链接kafka，创建分区消费者
	//2.2每个分区的消费者分别取出数据，通过SendToEs（）将数据发往ES
	err = kafka.Init([]string{cfg.KafkaCfg.Address}, cfg.KafkaCfg.Topic)
	if err != nil {
		fmt.Printf("init kafka consumer failed,err:%v\n", err)
		return
	}
	//3.发往es
	select{}
}
