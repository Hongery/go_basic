package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //调用init()
)

//Go 事物
var db *sql.DB //是一个连接池对象
type user struct {
	id int
	name string
	age string
}

//初始化数据库连接
func initDB() (err error){
	//数据库信息  "用户名：密码@tcp(ip地址:端口)/数据库名称"
	dsn := "root:root@tcp(127.0.0.1:3306)/sql_test"
	//链接数据库
	db, err = sql.Open("mysql", dsn) //不会校验用户名和密码是否正确
	if err != nil {
		fmt.Printf("dsn :%s invalid,err :%v\n", dsn, err)
		return
	}
	err = db.Ping() //尝试连接数据库
	if err != nil {
		fmt.Printf("open%s failed,err:%v\n", dsn, err)
		return
	}
	//设置数据库连接池的最大连接数
	db.SetMaxOpenConns(10)
	return
}

//事物操作示例
func transactionDemo(){
	tx, err := db.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	sqlStr1 := "Update user set age=30 where id=?"
	_, err = tx.Exec(sqlStr1, 2)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql1 failed, err:%v\n", err)
		return
	}
	sqlStr2 := "Update user set age=40 where id=?"
	_, err = tx.Exec(sqlStr2, 4)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	err = tx.Commit() // 提交事务
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("commit failed, err:%v\n", err)
		return
	}
	fmt.Println("exec trans success!")
}

func main(){
	err :=initDB()
	if err != nil {
		fmt.Printf("init DB failed,err :%v\n",err)
		return
	}
	fmt.Println("连接数据库成功" )
	transactionDemo()
}



