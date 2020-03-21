package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //调用init()
)

//Go链接Mysql示例
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

//查询单个记录
func queryOne(id int){
	var u1 user
	//1.写查询单条记录的sql语句
	sqlStr :=`select id,name,age from user where id=?;`
	//2.执行
	rowObj :=db.QueryRow(sqlStr,id)//从连接池里那一个连接出来，去数据库查询
	/*for i:=0;i<10;i++{ //测试最大连接数
		db.QueryRow(sqlStr,2)
	}*/
	//3.拿到结果
	rowObj.Scan(&u1.id,&u1.name,&u1.age)//必须对rowObj对象调用Scan方法，因为该方法会释放数据库连接
	/*//执行并拿到结果
	//db.QueryRow(sqlStr,2).Scan(&u1.id,&u1.name,&u1.age)
	*/
	//打印结果
	fmt.Printf("u1:%#v\n",u1)
}

//查询多条数据
func queryMore(id int){
	sqlStr := `select id,name,age from user where id > ?`
	rows,err := db.Query(sqlStr,id)
	if err != nil {
		fmt.Printf("query failed ,err :%v",err)
		return
	}
	//非常重要：关闭rows释放持有的数据库连接
	defer rows.Close()

	//循环读取结果集中的数据
	for rows.Next(){
		var u user
		err :=rows.Scan(&u.id,&u.name,&u.age)
		if err !=nil{
			fmt.Printf("scan failed.err :%v\n",err)
			return
		}
		fmt.Printf("id :%d name:%s age:%s\n",u.id,u.name,u.age)
	}
}

//插入数据 可以传入参数
func insert(){
	sqlStr := `insert into user(name,age) values(?,?)`
	ret,err :=db.Exec(sqlStr,"wa",11)
	if err != nil{
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	//如果是插入数据的操作，能够拿到插入数据的id
	id,err :=ret.LastInsertId()
	if err != nil{
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", id)
}

//更新数据
func updata(){
	sqlStr :=`update user set age=? where id=?;`
	ret,err :=db.Exec(sqlStr,39,2)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

//删除数据
func delete(){
	sqlStr :="delete from user where id=?"
	ret,err :=db.Exec(sqlStr,3)
	if err != nil {
		fmt.Printf("delete failed,err :%v\n",err)
		return
	}
	n,err := ret.RowsAffected()//操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffcted failed,err :%v\n",err)
		return
	}
	fmt.Printf("delete success ,affected rows:%d\n",n)
}

//预处理插入
func prepareInsert(){
	sqlStr:="insert into user(name,age) values (?,?)"
	stmt,err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed , err :%v\n ",err)
		return
	}
	defer stmt.Close()

	/*_,err := stmt.Exec("halo",19)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}*/
	//后续只需要拿到stmt去执行一些操作
	var m = map[string]int{
		"jack":20,
		"liuhuang":10,
		"liuzhen":10,
	}
	for k,v := range m{
		stmt.Exec(k,v)
	}
	fmt.Println("insert success.")
}

func main() {
	err :=initDB()
	if err != nil {
		fmt.Println()
	}
	fmt.Println("连接数据库成功！")
	queryOne(2)
	queryMore(0)
	//insert()
	//updata()
	// delete()
	prepareInsert()
	queryMore(0)
}
