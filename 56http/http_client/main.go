package main

import (
	"net/http"
	"fmt"
	"net/url"
	"io/ioutil"
)
//共用一个client使用于 请求比较频繁
var (
	client = http.Client{
	Transport:&http.Transport{
		DisableKeepAlives:false,
	},
}
)

func main(){
	/*resp , err := http.Get("http://127.0.0.1:9090/xxx/?name周琳&age=18")
	if err != nil {
		fmt.Printf("get url failed,err:%v",err)
		return
	} */
	data :=url.Values{}
	urlObj,_:=url.Parse("http://127.0.0.1:9090/xxx/")
	data.Set("name","周琳")
	data.Set("age","18")
	queryStr :=data.Encode() //URL encode编码后的url
	fmt.Println(queryStr)
	urlObj.RawQuery = queryStr
	req,err := http.NewRequest("GET",urlObj.String(),nil)
	//请求不是特别频繁，用完就关闭该链接
	//禁用keepAlive的client
	/*tr :=&http.Transport{
		DisableKeepAlives:true,
	}
	client := http.Client{
		Transport:tr,
	}*/
	resp,err :=client.Do(req)
	//resp,err :=http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("get url failed,err:%v",err)
		return
	}
	defer resp.Body.Close()//一定要记得关闭连接
	//发请求
	//resp中把服务端返回的数据读出来
	//var data []byte
	//resp.Body.Read()
	//resp.Body.Close()
	b,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read resp.Body failed,err:%v",err)
		return
	}
	fmt.Println(string(b))

}
