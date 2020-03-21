package main

import (
	"fmt"
	"os"
)

//1.文件对象类型
//2.获取文件对象的详细信息
 func main(){
	 fileObj,err :=os.Open("./main.go")
	 if err != nil {
		 fmt.Printf("file open failed err:%v",err)
		 return 
	 }
	 //1.文件对象类型
	 fmt.Printf("%T\n",fileObj)
	 //2.获取对象的详细信息
	 fileInfo,err :=fileObj.Stat()
	 if err != nil {
		fmt.Printf("file open failed err:%v",err)
		return 
	 }
	 fmt.Printf("文件大小是：%dB\n",fileInfo.Size())//文件大小509B

 }