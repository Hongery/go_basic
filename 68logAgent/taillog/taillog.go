package taillog

import (
	"studygo/day01/68logAgent/kafka"
	//"studygo/day01/68logAgent/etcd"
	"fmt"

	"github.com/hpcloud/tail"
	"context"
)

//专门从日志文件手机及日志的模块

var (
	tailObj *tail.Tail
)

//TailTask 一个日志收集的任务
type TailTask struct {
	Path     string
	Topic    string
	Instance *tail.Tail
	//为了实现推出t.run()
	ctx context.Context
	cancelFunc context.CancelFunc
}

//NewTailTask 初始化结构体
func NewTailTask(path, topic string) (tailObj *TailTask) {
	ctx,cancel:=context.WithCancel(context.Background())
	tailObj = &TailTask{
		Path:  path,
		Topic: topic,
		ctx:ctx,
		cancelFunc:cancel,
	}
	tailObj.Init() //根据路径去打开对应的日志
	return
}

//Init 初始化配置
func (t *TailTask) Init() {
	config := tail.Config{
		ReOpen:    true,                                 //重新打开
		Follow:    true,                                 //是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件的哪个地方开始读
		MustExist: false,                                //文件不存在不报错
		Poll:      true,
	}
	var err error
	tailObj, err = tail.TailFile(t.Path, config)
	if err != nil {
		fmt.Println("tail file failed ,err :", err)
	}

	go t.Run() //开启后台的一个goroutine，采集日志发送到kafka

}

//Run a
func (t *TailTask) Run() {
	for {
		select {
		case <-t.ctx.Done():
			fmt.Printf("tail_taskL:%s_%s,  over\n",t.Path,t.Topic)
			return
		case line := <-t.ReadChan():
			//3.2发往kafka
			kafka.SendToChan(t.Topic, line.Text)
		}
	}
}

//ReadChan 通道中读取信息
func (t *TailTask) ReadChan() <-chan *tail.Line {
	return tailObj.Lines
}
