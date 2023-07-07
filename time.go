package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	   用sleep实现定时器
	*/
	fmt.Println("start")
	tiker := time.NewTicker(time.Second)
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		<-tiker.C
	}

}

//定时器
func test()  {
	ticker := time.NewTicker(time.Second * 1) // 运行时长
	ch := make(chan int)
	go func() {
		var x int
		for x < 10 {
			select {
			case <-ticker.C:
				x++
				fmt.Printf("%d\n", x)
			}
		}
		ticker.Stop()
		ch <- 0
	}()
	<-ch // 通过通道阻塞，让任务可以执行完指定的次数。

	fmt.Printf("%s\n", "end")
}