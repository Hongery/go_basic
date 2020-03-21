package main

import (
	"os"
	"fmt"
	"bufio"
	"io/ioutil"
)

//打开文件写内容
func demo1(){
	//创建文件   添加内容     清空  三个都执行，会先将原先的内容清空
	file,err :=os.OpenFile("xx.txt",os.O_CREATE|os.O_APPEND|os.O_TRUNC,0666)
	if err !=nil{
		fmt.Printf("open file failed err :%v\n",err)
		return
	}
	//write 字节
	file.Write([]byte("zhoulin menbi le\n"))
	//writestring 字符串
	file.WriteString("zhoulinhahahah")

	defer file.Close()
}

func demo2(){
	//创建文件   添加内容  清空
	file,err :=os.OpenFile("xx.txt",os.O_CREATE|os.O_APPEND|os.O_TRUNC,0666)
	if err !=nil{
		fmt.Printf("open file failed err :%v\n",err)
		return
	}
	defer file.Close()
	//创建一个写的对象
	wr :=bufio.NewWriter(file)
	wr.WriteString("heloworld\n") //写到缓存中
	wr.Flush()  //将缓存的内容写入到文件当中
}

//直接往文件中写东西
func demo3(){
	str := "hello 沙河"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}

func  main(){
	//demo1()
	//demo2()
	demo3()
}


