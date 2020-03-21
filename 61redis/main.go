package main

import (
	"github.com/go-redis/redis"
	"fmt"
)

//redis

// var redisdb *redis.Client 

// func initRedis()(err error){
// 	redisdb = redis.NewClient(&redis.Options{
// 		Addr:"127.0.0.1:6379",
// 		Password:"", //no password set
// 		DB:0,       //use default DB
// 	})
// 	_,err =redisdb.Ping().Result()
// 	if err != nil {
// 		fmt.Println("failed connect",err)
// 		return
// 	}
// 	return nil
// }

func ExampleNewClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}
func main(){
	// err :=initRedis()
	// if err !=nil{
	// 	fmt.Println("err",err)
	// 	return
	// }
	ExampleNewClient()
	fmt.Println("succeses")
}
