package main

import (
	"fmt"
	"os"
)
/*
	函数版学生管理系统，
	写一个系统能够产看，新增学生，删除学生
*/
var (
	allStudent map[int64]*student //变量声明
)
type student struct {
	id   int64
	name string
}
func newStudnet(id int64,name string) *student{
	return &student{
		id : id ,
		name : name,
	}
}
//展示所有学生
func showallStudent() {
	for k,v :=range allStudent{
		fmt.Printf("学号：%d 姓名：%s\n",k,v.name,)
	}
}
//添加学生
func addStudent() {
	//创建一个学生
	var (
		id int64
		name string
	)
	fmt.Println("请输入学生学号：")
	fmt.Scanln(&id)
	fmt.Println("请输入学生姓名：")
	fmt.Scanln(&name)
	newStu :=newStudnet(id,name)
	//2追加到allstudent这个map中
	allStudent[id]=newStu
}
func deleteStudent() {
	//1.请输入删除学生的学号
	var deleteId  int64
	
	fmt.Println("请输入你要删除学生的学号:")
	fmt.Scanln(&deleteId)
	//2.去allStudent这个map中根据学号删除对应的键值对
	for k,_:=range allStudent{
		if k==deleteId{
			delete(allStudent,deleteId)
			fmt.Println("删除成功")
		}
	}
}

func main() {
	allStudent =make(map[int64]*student,48) //初始化（开辟内存空间）
	for{
	//1.打印菜单
	fmt.Println("欢迎学生管理系统")
	fmt.Println(`
		1.查看所用学生
		2.新增学生
		3.删除学生
		4.退出
	`)
	fmt.Println("请输入你要干啥:")
	//2.等待用户选择要做什么
	var choice int
	fmt.Scanln(&choice)
	fmt.Printf("你选择了%d这个选项\n", choice)
	//3.执行对应的函数
	switch choice {
	case 1:
		showallStudent()
	case 2:
		addStudent()
	case 3:
		deleteStudent()
	case 4:
		os.Exit(1)
	default:
		fmt.Println("gun~")
	}
}
}
