package main

import (
	//"github.com/hpcloud/tail"
	"fmt"
	"studygo/day01/68logAgent/conf"
	"studygo/day01/68logAgent/etcd"
	"studygo/day01/68logAgent/kafka"
	"studygo/day01/68logAgent/taillog"
	"studygo/day01/68logAgent/utils"
	"time"

	"sync"

	"gopkg.in/ini.v1"
)

//logAgent 入口函数

var (
	cfg = new(conf.AppConf)
)

/*
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
*/
//读取信息之前先使用67etcd_put.exe
func main() {
	//0.加载配置文件
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("load ini failed ,err :%v", err)
		return
	}
	//1. 初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.ChanMaxSize)
	if err != nil {
		fmt.Println("init kafka failed", err)
		return
	}
	fmt.Println("kafka success")

	//2 初始化etcd
	//5*time.Second
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Println("init etcd failed", err)
		return
	}
	fmt.Println("etcd success")

	//为了实现每个logagent都拉去自己独有的配置信息，所以要以自己的Ip地址作为区分
	ipStr, err := utils.GetOutbouneIP() //本地IP192.168.0.104,
	if err != nil {
		panic(err)
	}
	etcdConfKey := fmt.Sprintf(cfg.EtcdConf.Key, ipStr)
	//2.1 从etcd中获取日志收集项的配置信息
	logEntryConf, err := etcd.GetConf(etcdConfKey)
	//logEntryConf, err := etcd.GetConf(cfg.EtcdConf.Key)
	if err != nil {
		fmt.Printf("etcd.GetConf failed ,err %v\n", err)
		return
	}
	fmt.Printf("get conf from etcd success,%v\n", logEntryConf)
	//2.2 派一个哨兵去监视日志收集项的变化（有变化及时通知我的logAgent实现加载配置）
	//newConfChan :=taillog.NewConfChan()//从taillog中取出对外暴露的通道
	for index, value := range logEntryConf {
		fmt.Printf("index:%v value :%v\n", index, value)
	}
	//3.收集日志发往kafka
	taillog.Init(logEntryConf)
	newConfChan := taillog.NewConfChan() //从taillog中取出对外暴露的通道
	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.WatchConf(etcdConfKey, newConfChan)
	wg.Wait()

}
