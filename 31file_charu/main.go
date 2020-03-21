package main

import (
	"fmt"
	"os"
)
func f2(){
	fileObj,err :=os.OpenFile("./sb.txt",os.O_RDWR,0644)
	if err !=nil {
		fmt.Printf("open file failed err:%v",err)
		return
	}
	defer fileObj.Close()
	//光标移动到b
	fileObj.Seek(3,0) //txt中会有\r \n 两个字符
	var ret [1]byte
	n,err :=fileObj.Read(ret[:])
	if err !=nil {
		fmt.Printf("open file failed err:%v",err)
		return
	}
	fmt.Println(string(ret[:n]))
}


func f3(){
	//打开要操作的文件(读写)
	fileObj,err :=os.OpenFile("./sb.txt",os.O_RDWR,0633)
	if err !=nil {
		fmt.Printf("open file failed err:%v",err)
		return
	}
	//因为没有办法直接在文件中间插入内容，所以要借助一个临时文件
	tmpFile,err :=os.OpenFile("./sb.tmp",os.O_CREATE|os.O_TRUNC|os.O_WRONLY,0644)
	if err !=nil{
		fmt.Printf("open file failed err:%v",err)
		return
	}
	defer tmpFile.Close()

	//读取源文件写入临时文件
	var ret [1]byte
	n,err :=fileObj.Read(ret[:])
	if err!=nil {
		fmt.Printf("open file failed err:%v",err)
		return
	}
	//写入临时文件
	tmpFile.Write(ret[:n])
	//再写入要插入的内容
	var s []byte
	s =[]byte{'c'}
	tmpFile.Write(s)
	//紧接着把源文件后续的内容写入临时文件
	var x [1024]byte
	for {
		n,err :=fileObj.Read(x[:])
		if err !=nil {
			tmpFile.Write(x[:n])
			break
		}
		if err !=nil{
			fmt.Printf("open file failed err:%v",err)
			return
		}
		tmpFile.Write(x[:n])
	}
	//源文件后续的也写入了临时文件
	fileObj.Close()
	tmpFile.Close()
	os.Rename("./sb.tmp","./sb.txt")
}
func main(){
	f3()
}