package main

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

func main() {
	// 创建Redis连接
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for {

		currentTime := time.Now()
		fmt.Println("当前时间：", currentTime)

		nowDate := fmt.Sprintf("当前时间： %d-%d-%d %d:%d:%d", currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour(), currentTime.Minute(), currentTime.Second())

		sendMsg := "Hello World : " + nowDate

		// 获取redis信息

		result, err := redis.Values(conn.Do("lpop", "key")) //读
		if err != nil {
			panic(err)
		}
		fmt.Println("queue长度", result)

		// 发送消息
		_, err = conn.Do("LPUSH", "myqueues", sendMsg)
		if err != nil {
			panic(err)
		}

		//sleep 1秒
		time.Sleep(time.Duration(1) * time.Second)

	}

}
