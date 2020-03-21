package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

//UDP client

func main() {
	//拨号建立连接
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40,
		Zone: "",
	})
	if err != nil {
		fmt.Println("lianjie failed ,err :", err)
		return
	}
	defer socket.Close()
	var reply [1024]byte
	reader := bufio.NewReader(os.Stdin) //获取读取对象
	for {
		fmt.Println("请输入内容：")
		msg, _ := reader.ReadString('\n')
		socket.Write([]byte(msg))
		//收到回复数据
		n, _, err := socket.ReadFromUDP(reply[:])
		if err != nil {
			fmt.Println("redv reply msg failed ,err:", err)
			return
		}
		fmt.Println("收到下信息回复：", string(reply[:n])) //准换为大写
	}

	/*sendData:=[]byte("hello server")//发送数据
	_, err =socket.Write(sendData)
	if err !=nil {
		fmt.Println("send data failed,err",err)
		return
	}
	data := make([]byte,4096)
	n,remoteAddr,err := socket.ReadFromUDP(data)//接受数据
	if err != nil {
		fmt.Println("recept failed ,err",err)
		return
	}
	fmt.Printf("recv:%v addr:%v count:%v \n",string(data[:n]),remoteAddr,n)*/
}
