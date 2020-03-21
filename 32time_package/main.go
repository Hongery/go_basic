package main

import (
	"fmt"
	"time"
)
//格式化时间
func geshihua(){
	//把语言中时间对象 转换成字符串类
	now :=time.Now()
	fmt.Println(now.Format("2016-01-02"))
	fmt.Println(now.Format("2016-10-01 15:04:02"))
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}

//定时器
func dingshiqi(){
	//定时器
	timer :=time.Tick(time.Second) //一秒钟执行一次，可以换成Hour小时
	for t :=range timer {
		fmt.Println(t)
		//2020-01-30 16:56:24.5012857 +0800 CST m=+1.009000401
	}
}

//解析字符串格式的时间
func jiexi(){
	now := time.Now()
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2019/08/04 14:15:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)//2019-08-04 14:15:20 +0800 CST
	fmt.Println(timeObj.Sub(now))//-4301h6m21.3074446s
}

//Sub两个时间差
func shijiancha(){
	now :=time.Now()
	nextYear,err :=time.Parse("2006-01-02","2020-01-29")
	fmt.Println(nextYear)//2020-01-29 00:00:00 +0000 UTC
	if err !=nil{
		fmt.Println(err)
		return
	}
	//now-nextYear
	d := now.Sub(nextYear)
	fmt.Println(d)//35h49m26.6250517s

}
//时区
func f2(){
	now :=time.Now()
	fmt.Println(now)
	//明天这个时间
	//按照指定格式解析一个字符串格式的时间
	time.Parse("2006-01-02 15:03:05","2020-01-30 14:41:50")
	//按照东八区的失去和格式取解析一个字符串格式的时间
	//根据字符串加载时区
	loc,err :=time.LoadLocation("Asia/Shanghai")
	if err !=nil{
		fmt.Println(err)
		return
	}
	//根据制定时区解析时间
	timeObj,err :=time.ParseInLocation("2006-01-02 15:03:05","2020-01-30 14:41:50",loc)
	if err !=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)

	//时间对象相减
	td :=now.Sub(timeObj)
	fmt.Println(td)
}


func main(){
	now := time.Now() //获取当前时间
	fmt.Println(now) //2020-01-30 16:32:27.7441079 +0800 CST m=+0.001499301
	fmt.Println(now.Year())//2020
	fmt.Println(now.Month())//1
	fmt.Println(now.Day())//30
	fmt.Println(now.Hour())//16
	fmt.Println(now.Minute())//32
	fmt.Println(now.Second())//47
	//时间戳
	fmt.Println(now.Unix())//1580373589
	fmt.Println(now.UnixNano())//纳秒时间戳 1580373589465713100
	//将时间戳转化为时间格式
	ret :=time.Unix(1580373589,0)
	fmt.Println(ret) //2020-01-30 16:39:49 +0800 CST
	fmt.Println(ret.Year())//2020
	fmt.Println(ret.Day())//30

	//时间间隔
	fmt.Println(time.Second) //1s
	//now+24小时 1：30 16：56
	fmt.Println(now.Add(24*time.Hour)) //2020-01-31 16:56:23.4932978 +0800 CST m=+86400.001012501
	
	//geshihua()//格式化
	//dingshiqi()//定时器

	//按照对应的格式解析字符串类型的时间
	timeObj ,err :=time.Parse("2006-01-02","2020-01-30")
	if err !=nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)	//2020-01-30 00:00:00 +0000 UTC
	fmt.Println(timeObj.Unix()) //1580342400

	//jiexi()
	shijiancha()

	//sleep
	n :=5
	fmt.Println("开始sleep")
	time.Sleep(time.Duration(n)*time.Second)
	fmt.Println("5秒过去了")

	f2()
}
