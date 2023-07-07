package main

import (
	"fmt"
	"os"
)

func getDBC() (string, bool) {
	conn := os.Getenv("dbc")
	if conn == "" {
		return conn, false
	}
	return conn, true
}

func main() {
	a := getDBC()
	fmt.Println(a)
}
