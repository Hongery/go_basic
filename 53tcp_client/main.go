package main

import (
	"net"
	"fmt"
	"os"
	"bufio"
	"strings"
)

//tcp client

func main(){
	//1.与server端建立连接
	conn,err := net.Dial("tcp","127.0.0.1:20") //Dial拨号连接到指定网络上的地址。
	if err != nil {
		fmt.Println("dial 127.0.0.1:20 failed,err :",err)
		return
	}
	//2.发送数据
	/*go build 执行tcp_client.exe （msg）参数
	var msg string
	if len(os.Args)<2{
		msg="hello huanggang"
	}else {
		msg=os.Args[1]
	}*/
	reader :=  bufio.NewReader(os.Stdin) //获得reader对象
	for {
		fmt.Print("请说话：")
		msg,_:=reader.ReadString('\n')//读取直到输入换行符'\n'结束读取字符串
		msg =strings.TrimSpace(msg)//返回字符串，删除开头和结尾的空格
		if msg =="exit"{
			break
		}
		conn.Write([]byte(msg))//往服务端写数据
	}
	defer conn.Close()
}
