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

    // 订阅频道
    subConn := redis.PubSubConn{Conn: conn}
    subConn.Subscribe("mychannel")

    // 接收消息
    for {
        switch v := subConn.Receive().(type) {
        case redis.Message:
            fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
        case redis.Subscription:
            fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
        case error:
            panic(v)
        }
    }
}