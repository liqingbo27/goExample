package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	var version string

	err2 := db.QueryRow("SELECT VERSION()").Scan(&version)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(version)

	fmt.Println("successfully create table")

	mid := 8

	result, err2 := db.Query("SELECT * FROM user WHERE id = ?", mid)
	if err2 != nil {
		log.Fatal(err2)
	}

	for result.Next() {

		var id int
		var name string

		err = result.Scan(&id, &name)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("id: %d, name: %s\n", id, name)
	}

}
