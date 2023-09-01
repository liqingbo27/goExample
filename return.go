package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("start")

	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
		}
	}()

	for {

		a()

		b()

		c()

		// sleep 1秒
		time.Sleep(time.Duration(3) * time.Second)
	}

}

func a() {
	fmt.Println("a")
}

func b() {
	panic("bbb")
	fmt.Println("b")
}

func c() {
	fmt.Println("c")
	return
}
