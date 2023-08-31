package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	// 创建Redis连接
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("start queue for")

	// 接收消息
	for {
		reply, err := conn.Do("BRPOP", "myqueues", 0)
		if err != nil {
			panic(err)
		}
		messages, err := redis.Strings(reply, err)
		if err != nil {
			panic(err)
		}
		fmt.Printf("message: %s\n", messages[1])
	}
}
