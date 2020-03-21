package main

import (
	"github.com/olivere/elastic/v7"
	"fmt"
	"context"
)

//ES insert data demo
type Student struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Married bool `json:"married"`
}

func main() {
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Println("connect to es success")
	p1 := Student{Name: "rion", Age: 22, Married: false}
	// 链式调用
	put1, err := client.Index().
		Index("student").
		BodyJson(p1).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed student %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}