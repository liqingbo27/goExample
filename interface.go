package main

import (
	"fmt"
)

func main() {
	test1()
	fmt.Println()
	fmt.Println("-------------------------")
	test2()
}

func test1() {
	mapInterface := make(map[interface{}]interface{})
	mapString := make(map[string]string)

	mapInterface["k1"] = 1
	mapInterface[3] = "hello"
	mapInterface["world"] = 1.05

	for key, value := range mapInterface {
		strKey := fmt.Sprintf("%v", key)
		strValue := fmt.Sprintf("%v", value)

		mapString[strKey] = strValue
	}

	fmt.Printf("%#v", mapString)
}

func test2() {
	m := make(map[string]interface{})
	m["int"] = 123
	m["string"] = "hello"
	m["bool"] = true

	for _, v := range m {
		switch v.(type) {
		case string:
			fmt.Println(v, "is string")
		case int:
			fmt.Println(v, "is int")
		default:
			fmt.Println(v, "is other")
		}
	}
	fmt.Println(m)
}
