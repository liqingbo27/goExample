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

	if _, err = conn.Do("SET", "foo", "liqingbo"); err != nil {
		fmt.Println(err)
		return
	}

	exists, err := redis.Bool(conn.Do("EXISTS", "foo"))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v\n", exists)

	ginfo, err := conn.Do("GET", "foo")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 获取元素
	element := string(ginfo.([]byte))

	fmt.Println(element)

}
