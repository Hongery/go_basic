package main

import (
	"context"
	"fmt"
	"time"
	"go.etcd.io/etcd/clientv3"
)


func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()
	//watch
	//派一个哨兵一直监视huanggang这个key的变化（新增、修改、删除）
	ch :=cli.Watch(context.Background(),"huanggang") //<-chan WatchResponse
	//从通道尝试取值
	for wresp :=range ch{
		for _, evt := range wresp.Events{
			//fmt.Printf("Type :%v key:%v value :%v\n",evt,evt.Kv.Key,evt.Kv.value)
		fmt.Println(evt)
			}
	}
}
