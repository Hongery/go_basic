package taillog

import (
	"studygo/day01/68logAgent/etcd"
	"fmt"
	"time"
)

var tskMgr *tailLogMgr

type tailLogMgr struct{
	 logEntry []*etcd.LogEntry
	 tskMap map[string]*TailTask
	 newConfChan chan []*etcd.LogEntry
}
//Init
func Init(logEntryConf []*etcd.LogEntry){
	tskMgr = &tailLogMgr{
		logEntry:logEntryConf, //把当前的日志收集项配置信息保存起来
		tskMap:make(map[string]*TailTask,16),
		newConfChan:make(chan []*etcd.LogEntry),//无缓冲区通道
	}
	for _,logEntry := range logEntryConf{
		//conf :*etcd.LogEntry  logEntry.Path:要收集的日志文件的路径
		//初始化的时候起了多少个tailtask都要记下来，为了后续的判断方便
		tailObj :=NewTailTask(logEntry.Path,logEntry.Topic)//TailTask struct
		mk :=fmt.Sprintf("%s_%s",logEntry.Path,logEntry.Topic)//拼接路径和topic ，路径唯一
		tskMgr.tskMap[mk]=tailObj
	}
	go tskMgr.run()

}


//监听自己的newConfChan 有了新的配置之后就做新的处理
func (t *tailLogMgr)run(){
	for {
		select {
		case newConf :=<- t.newConfChan:
			for _,conf := range newConf{
				mk :=fmt.Sprintf("%s_%s",conf.Path,conf.Topic)//拼接路径和topic ，路径唯一
				_,ok:=t.tskMap[mk]
				if ok {
					//原来就有不需要操作
					continue
				}else {
					//新增的
					tailObj:=NewTailTask(conf.Path,conf.Topic)
					t.tskMap[mk]=tailObj
				}

			}
			for _,c1:=range t.logEntry{//从原来的配置中一次拿出配置项
				isDelete :=true
				for _,c2:=range newConf{//去新的配置项中去比较
				if c2.Path==c1.Path && c2.Topic==c1.Topic {
					isDelete = false
					continue
					}
				}
				if isDelete{
					//把c1对应的这个tailObj给停掉
					mk :=fmt.Sprintf("%s_%s",c1.Path+c1.Topic)
					t.tskMap[mk].cancelFunc()
				}
			}
			//1.配置新增
			//2.配置删除
			//3.配置变更
			fmt.Println("新的配置起来了",newConf)
		default:
			time.Sleep(time.Second)
		}
	}
}

//一个函数，向外暴露tskMgr的newConfChan
func NewConfChan()chan <- []*etcd.LogEntry{//只读
	return tskMgr.newConfChan
}
