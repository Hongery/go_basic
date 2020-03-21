package main

import (
	"os"
	"fmt"
)

//学生管理系统
var smr StudentMgr //声明一个全局的管理者变量


//菜单函数
func showMenu(){
	fmt.Println("welcome sms！")
	fmt.Println(`
	1.查看所有学生
	2.添加学生
	3.修改学生
	4.删除学生
	5.退出
	`)
}

func main(){
  smr = StudentMgr { //修改全局变量
	  AllStudent : make(map[int64]student,100),
	}
	for{
		
		showMenu()
		//等待用户输入选项
		fmt.Println("请输入序列号：")
		var choice int 
		fmt.Scanln(&choice)
		fmt.Println("您输入的是：",choice)
		switch choice {
		case 1:
			smr.showStudent()
		case 2:
			smr.addStudent()
		case 3:
			smr.editStudent()
		case 4:
			smr.deleteStudent()
		case 5:
			os.Exit(1)	
		default:
			fmt.Println("gun!!")
		}

	}
}
