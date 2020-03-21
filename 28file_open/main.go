package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
	"io/ioutil"
)

//打开文件  字节读取
func readfromFile(){
	fileObj, err := os.Open("./main.go")  //在VScode编译器中可移执行，在IDEA中报错open file failed err:open ./main.go: The system cannot find the file specified.
	if err != nil{
		fmt.Printf("open file failed err:%v",err)
		return
	}
	//记得关闭文件
	defer fileObj.Close()

	//读文件
	//var tmp [128]byte
	var tmp = make([]byte,128) //指定读的长度
	//循环读取文件中的数据，一次读取128个字节
	for {
		n, err := fileObj.Read(tmp[:])
		if err != nil {
			fmt.Printf("open file failed err: %v", err)
			return
		}
		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(tmp))
		if n < 128 {
			return
		}
	}

}

//利用bufio包读取文件  一行一行读取
func readFromFileByBufio(){
	fileObj, err := os.Open("./main.go")  //在VScode编译器中可移执行，在IDEA中报错open file failed err:open ./main.go: The system cannot find the file specified.
	if err != nil{
		fmt.Printf("open file failed err:%v",err)
		return
	}
	//记得关闭文件
	defer fileObj.Close()

	//创建一个用来从文件中读取内容的对象
	reader :=bufio.NewReader(fileObj)
	for{
		line ,err :=reader.ReadString('\n') //string err
		if err ==io.EOF{ //EOF 如果读取到文件的末尾，结束
			return
		}

		if err != nil{
			fmt.Printf("read file failed err :%v",err)
			return
		}
		fmt.Print(line)
	}
}

//ioutil整个文件读取
func readFromFileByIoutil(){
	content ,err :=ioutil.ReadFile("./main.go")
	if err != nil{
		fmt.Printf("read file failed err :%v",err)
		return
	}
	fmt.Println(string(content))
}

func main(){
	//readFromFile()
	//readFromFileByBufio()
	readFromFileByIoutil()

}
