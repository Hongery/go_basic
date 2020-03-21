package main

import (
	"fmt"
	"net"
	"strings"
)

//UDP server

func main() {
	//获取连接对象
	socket, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40,
		Zone: "",
	})
	if err != nil {
		fmt.Println("listen UDP failed,err:", err)
		return
	}
	defer socket.Close()
	//不需要像tcp建立连接，直接发送数据
	var data [1024]byte
	for {
		n, addr, err := socket.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read from UDP ,err", err)
			return
		}
		fmt.Println(data[:n])
		reply := strings.ToUpper(string(data[:n]))
		//发送数据
		socket.WriteToUDP([]byte(reply), addr)
	}
}
