package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

//net/http/server

func f(w http.ResponseWriter,r *http.Request){
	b,err :=ioutil.ReadFile("xx.html")
	if err !=nil {
		w.Write([]byte(fmt.Sprintf("%v",err)))
		return
	}
	w.Write(b)

}

func f2(w http.ResponseWriter,r *http.Request){
	//对于GET请求，参数都放在URL上（query param），请求体中没有参数
	//fmt.Println(r.URL.Query()) // 使用r.URL 结果是‘/xxx/’ ，使用QUERY自动帮我们识别URL中的参数
	queryParam :=r.URL.Query()
	name := queryParam.Get("name")
	age :=queryParam.Get("age")
	fmt.Println(name,age)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))//我在服务端打印客户端发来的请求Body
	w.Write([]byte("ok"))
}

func main(){
	http.HandleFunc("/hello",f)
	http.HandleFunc("/xxx/",f2)
	http.ListenAndServe("127.0.0.1:9090",nil)
}
