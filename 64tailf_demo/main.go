package main

import (
	"github.com/hpcloud/tail"
	"fmt"
	"time"
)

//tail 测试

func main(){
	fileName :="./my.log"
	config :=tail.Config{
		ReOpen :true,//重新打开
		Follow:true,//是否跟随
		Location:&tail.SeekInfo{Offset:0,Whence:2},//从文件的哪个地方开始读
		MustExist:false,//文件不存在不报错
		Poll:true,
	}
	tails,err :=tail.TailFile(fileName,config)
	if err != nil {
		fmt.Println("tail file failed ,err :",err )
		return
	}
	var (
		line *tail.Line
		ok bool
	)
	for {
		line,ok =<-tails.Lines
		if !ok{
			fmt.Printf("tail ile close reOpen ,filename s%\n",tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("line:",line.Text)
	}

}
