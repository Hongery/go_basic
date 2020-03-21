package main


import (
	"fmt"
)

//学生
type student struct {
	id   int64
	name string
}

//学生管理者
type StudentMgr struct {
	AllStudent map[int64]student
}

//查看学生
func (s StudentMgr) showStudent() {
	for _, stu := range s.AllStudent {
		fmt.Printf("学号 ：%d 姓名：%s ", stu.id, stu.name)
	}
}

//添加学生
func (s StudentMgr) addStudent() {
	//1.根据用户输入的内容创建一个新的学生
	var (
		id   int64
		name string
	)
	//获取用户输入的内容
	fmt.Print("请输入学号：")
	fmt.Scanln(&id)
	fmt.Print("请输入姓名：")
	fmt.Scanln(&name)
	newStu := student{
		id:   id,
		name: name,
	}
	//2.把新的学生添加到s.AllStudent这个map中
	s.AllStudent[newStu.id] = newStu
	fmt.Println("添加成功")

}

//修改学生
func (s StudentMgr) editStudent() {
	//1.获取用户输入的学号
	var stuID int64
	fmt.Print("请输入学号：")
	fmt.Scanln(&stuID)
	//2.展示该学号对应的学生信息，如果没有查无
	stuObj, ok := s.AllStudent[stuID]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	fmt.Printf("你要修改的学生学号：%d 姓名：%s", stuObj.id, stuObj.name)
	//3.修改输入的学生名字
	fmt.Println("请输入学生的新名字")
	var newName string
	fmt.Scanln(&newName)

	//4.更新学生的姓名
	stuObj.name = newName
	s.AllStudent[stuID] = stuObj

}

//删除学生
func (s StudentMgr) deleteStudent() {
	//1.请输入要删除的学生
	var stuID int64
	fmt.Print("请输入学号：")
	fmt.Scanln(&stuID)
	//2.去map中查找
	_, ok := s.AllStudent[stuID]
	if !ok {
		fmt.Println("查无此人")
		return
	}

	//3.有的话就删除，如何从map中删除键值对
	delete(s.AllStudent,stuID)
	fmt.Println("删除成功")
}
