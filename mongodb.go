package main

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
)

type User struct {
	Name     string    `bson:"name"`
	Age      int       `bson:"age"`
	Email    string    `bson:"email"`
	CreateAt time.Time `bson:"create_at"`
}

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	collection := session.DB("test").C("users")

	user := User{
		Name:     "Alice",
		Age:      28,
		Email:    "alice@example.com",
		CreateAt: time.Now(),
	}
	err = collection.Insert(user)
	if err != nil {
		panic(err)
	}

	var result []User
	err = collection.Find(nil).Sort("-create_at").Limit(10).All(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
