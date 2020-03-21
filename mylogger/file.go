package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

//往文件里面写日志相关代码
//main函数在mylogger_demo中
type FileLogger struct {
	Levle       LogLevel
	filePath    string //日志文件保存的路径
	fileName    string //日志文件保存的文件名
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
}

//NewFileLogger 构造函数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	f1 := &FileLogger{
		logLevel,
		fp,
		fn,
		nil,
		nil,
		maxSize,
	}
	err = f1.initFile()
	if err != nil {
		panic(err)
	}
	return f1
}

//初始化fileObj ，errFileObj
func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	//hello.log 在35mylogger——demo
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open file failed err :%v\n", err)
		return err
	}
	//hello.log.err
	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open file failed err :%v\n", err)
		return err
	}
	//日志文件都已经打开
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}

//判断是否需要记录该文件
func (f *FileLogger) enable(logLevel LogLevel) bool {
	return f.Levle <= logLevel //>=INFO级别的都输出，可以自己设置
}

//对文件大小进行切割，判断是否需要切割
func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("file open failed err:%v", err)
		return false
	}
	//如果当文件大小 大于等于 日志文件的最大值 就应该返回true
	return fileInfo.Size() >= f.maxFileSize //maxFileSize 在构造函数NewFileLogger中自己设置
}

//切割文件
func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	//需要切割日志文件
	nowStr := time.Now().Format("20060102150405000")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v", err)
		return nil, err
	}
	logName := path.Join(f.filePath, fileInfo.Name()) //名字拼接
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr)
	//1.关闭当前的日志文件
	file.Close()
	//2.备份一下rename xx.log -> xx.log.bak201908031709
	os.Rename(logName, newLogName) //文件名替换
	//3.打开一个新的日志文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("file err:%v", err)
		return nil, err
	}
	//4.将打开的新日志文件对象复制给 f.fileObj
	// f.fileObj=fileObj
	return fileObj, nil
}

//输出函数名，函数文件，行号，记录日志的方法
func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3) //回到函数调用的第三层
		if f.checkSize(f.fileObj) {
			//需要切割日志文件
			newFile, err := f.splitFile(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		fmt.Fprintf(f.fileObj, "[%s][%s] [%s:%s:%d]%s\n", now.Format("2006-09-02 15:04:09"), getLogString(lv), fileName, funcName, lineNo, msg)
		if lv > ERROR {
			if f.checkSize(f.errFileObj) {
				//需要切割日志文件
				newFile, err := f.splitFile(f.errFileObj)
				if err != nil {
					return
				}
				f.errFileObj = newFile
			}
			//如果要记录的日志大于等于ERROR级别，我还要在err日志文件中再记录一遍
			fmt.Fprintf(f.fileObj, "[%s][%s] [%s:%s:%d]%s\n", now.Format("2006-09-02 15:04:09"), getLogString(lv), fileName, funcName, lineNo, msg)
		}
	}
}

func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}
func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}
func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}
func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}
func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}

func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
