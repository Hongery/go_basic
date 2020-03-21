package main

import (
	"flag"
	"fmt"
	"time"
)

//flag 获取名行参数
func main() {
	//创建标志位
	// $ ./59flag_demo.exe -name=黄刚 -age=18 -married=true -ct=1000h
	name := flag.String("name", "王刚", "请输入姓名") //王刚是默认值，返回是指针
	age := flag.Int("age", 9000, "请输入真实年龄")
	married := flag.Bool("married", false, "结婚了")
	cTime := flag.Duration("ct", time.Second, "结婚了")

	//var name string
	//flag.StringVar(&name, "name", "张三", "姓名")  //使用var 返回的是值变量
	//开启使用  flag
	flag.Parse()
	//fmt.Println(name)//

	fmt.Println(*name)         //黄刚
	fmt.Println(*age)          //18
	fmt.Println(*married)      //true
	fmt.Println(*cTime)        //1000h0m0s  1000小时0分钟0秒
	fmt.Printf("%T\n", *cTime) //time.Duration

	fmt.Println(flag.Args())  //返回命令行参数后的其他参数，以[]string类型  //[]
	fmt.Println(flag.NArg())  //返回命令行参数后的其他参数个数  ,除了标志位以外的参数个数   //0
	fmt.Println(flag.NFlag()) //返回使用的命令行参数个数（使用标志位） 	//4

}
