package mylogger

import (
	"fmt"
	"time"
)

//往终端些日之相关内容

//Logger日志结构体
type ConsoleLogger struct {
	Level LogLevel
}

//NewLog 构造函数 返回level
func NewConsoleLogger(levelStr string) ConsoleLogger{
	level,err :=parseLogLevel(levelStr)
	if err !=nil {
		panic(err)
	}
	return ConsoleLogger{
		level,
	}
}

func (c ConsoleLogger)enable(logLevel LogLevel) bool{
	return c.Level<= logLevel //>=INFO级别的都输出，可以自己设置
}

//输出函数名，函数文件，行号
func (c ConsoleLogger)log(lv LogLevel,format string,a ...interface{}){
	if c.enable(lv){
	msg :=fmt.Sprintf(format,a...)
	now :=time.Now()
	funcName,fileName,lineNo :=getInfo(3)
	fmt.Printf("[%s][%s] [%s:%s:%d]%s\n", now.Format("2006-09-02 15:04:09"),getLogString(lv),fileName,funcName,lineNo,msg)
	}
}

func (c ConsoleLogger )Debug(format string,a ...interface{}){
	c.log(DEBUG,format,a...)
		
}
func (c ConsoleLogger )Info(format string,a ...interface{}){
	
		c.log(INFO,format,a...)	
}
func (c ConsoleLogger )Warning(format string,a ...interface{}){
		c.log(WARNING,format,a...)
}
func (c ConsoleLogger )Error(format string,a ...interface{}){
		c.log(ERROR,format,a...)	
}
func (c ConsoleLogger )Fatal(format string,a ...interface{}) {
		c.log(FATAL,format,a...)	
}

