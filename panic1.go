package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("start")

	for {

		defer func() { // 必须要先声明defer，否则不能捕获到panic异常
			if err := recover(); err != nil {
				fmt.Println(err) // 这里的err其实就是panic传入的内容，55
			}
		}()

		defer fmt.Printf("2\n")

		defer b()
		defer c()

		defer fmt.Printf("1\n")

		a()

		panic("在a这里中断") //下面的代码不执行了

		b()

		c()

		time.Sleep(time.Duration(1) * time.Second)
	}

}

func a() {
	fmt.Println("a")
}

func b() {
	fmt.Println("b")
}

func c() {
	fmt.Println("c")
	return
}
