package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
)

//专门往kafka写日志的模块

var (
	//声明一个全局连接kafka的生产者client
	client sarama.SyncProducer
)

//Init 初始化client
func Init(addr []string) (err error) {
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
	return
}

func SendToKafka(topic, data string) {
	//构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)
	//发送到kafka
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed ,err:", err)
		return
	}
	fmt.Printf("pid:%v offset :%v\n", pid, offset)
}
