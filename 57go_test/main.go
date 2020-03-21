package main

import (
	"fmt"
	"studygo/day01/57go_test/split"
)

func main(){
	res :=split.Split("abcdesfd","b")
	fmt.Println(res)
	result :=split.Splitstring("asdddfd","dd")
	fmt.Println(result)
}
