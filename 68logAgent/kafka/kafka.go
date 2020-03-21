package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
	"time"
)

//专门往kafka写日志的模块
type logData struct {
	topic string
	data string
}

var (
	//声明一个全局连接kafka的生产者client
	client sarama.SyncProducer
	logDataChan chan *logData
)

//Init 初始化client
func Init(addr []string,maxSize int) (err error) {
	config := sarama.NewConfig()
	//tailf包使用
	//发送完数据需要leader和follow都确认
	config.Producer.RequiredAcks = sarama.WaitForAll
	//新选出一个partition分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//成功操作的消息将在
	config.Producer.Return.Successes = true
	//连接kafka
	client, err = sarama.NewSyncProducer(addr, config)
	if err != nil {
		fmt.Println("producer closed err :", err)
		return
	}
	//初始化logDataChan
	logDataChan =make(chan *logData,maxSize)
	//开启后台的goroutine从通道中获取数据
	go SendToKafka()
	return
}
//把日志数据发送到内部的channel中
func SendToChan(topic ,data string){
	msg:=&logData{
		topic :topic,
		data:data,
	}
	logDataChan<-msg
}

//真正往kafka发送日志的函数
func SendToKafka() {
	for   {
		select {
		case lg:=<-logDataChan:

		//构造一个消息
		msg := &sarama.ProducerMessage{}
		msg.Topic = lg.topic
		msg.Value = sarama.StringEncoder(lg.data)
		//发送到kafka
		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("send msg failed ,err:", err)
			return
		}
		fmt.Printf("pid:%v offset :%v\n", pid, offset)
		default:
			time.Sleep(time.Millisecond*1000)
		}
	}
}
