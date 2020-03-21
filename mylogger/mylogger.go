package mylogger

import (
	"strings"
	"errors"
	"runtime"
	"fmt"
	"path"
	"time"
)

type LogLevel uint16

type Logger interface{
	Debug(format string,a ...interface{})
	Info(format string,a ...interface{})
	Error(format string,a ...interface{})
	Warning(format string,a ...interface{})
	Fatal(format string,a ...interface{})
}

const(
	UNKNOW LogLevel =iota //0往下累加
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)
func parseLogLevel(s string) (LogLevel,error){
	s =strings.ToLower(s)
	switch s{
	case "debug":
		return DEBUG , nil
	case "trace":
		return TRACE,nil
	case "info":
		return INFO,nil
	case "warning":
		return WARNING, nil
	case "error" :
		return ERROR , nil
	case "fatal":
		return FATAL , nil
	default:
		err :=errors.New("日志级别错误")
		return UNKNOW,err
	}
}
func getLogString(lv LogLevel) string{
	switch lv{
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}
	return "DEBUG"
}

//获取函数名，文件名，行号
func getInfo(skip int)(funcName,fileName string,lineNo int){
	pc,file,lineNo,ok :=runtime.Caller(skip)
	if !ok{
		fmt.Println("runtime.caller() failed")
		return
	}
	funcName =runtime.FuncForPC(pc).Name() //main.main
	fileName =path.Base(file)
	funcName=strings.Split(funcName,".")[1] //只留下main
	return
}
//输出函数名，函数文件，行号
func log(lv LogLevel,format string,a ...interface{}){
	msg :=fmt.Sprintf(format,a...)
	now :=time.Now()
	funcName,fileName,lineNo :=getInfo(3)
	fmt.Printf("[%s][%s] [%s:%s:%d]%s\n", now.Format("2006-09-02 15:04:09"),getLogString(lv),fileName,funcName,lineNo,msg)

}