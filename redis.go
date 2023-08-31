package main

import (
	"github.com/gomodule/redigo/redis"
)

func main() {
	// 创建Redis连接
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 发送消息
	_, err = conn.Do("PUBLISH", "mychannel", "Hello World!")
	if err != nil {
		panic(err)
	}
}
