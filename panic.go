package main

import "fmt"

func main() {

	fmt.Println("1")

	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("5")
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
		}
		fmt.Println("6")
	}()

	defer fmt.Println(4)
	defer fmt.Println(3)

	f()              //开始调用f
	fmt.Println("f") //这里开始下面代码不会再执行
}

func f() {
	fmt.Println("2")
	panic("异常信息")
	fmt.Println("7") //这里开始下面代码不会再执行
	fmt.Println("8")
}
