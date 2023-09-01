package main

import "fmt"

// Person 定义接口
type Person interface {
	GetName() string
	GetAge() uint32
}

// Student 定义类型
type Student struct {
	Name string
	Age  uint32
}

func (s Student) GetName() string {
	return s.Name
}
func (s Student) GetAge() uint32 {
	return s.Age
}

func main() {

	var student Student
	student.Age = 12
	student.Name = "小明"

	var person Person
	person = student
	fmt.Printf("name:%s,age: %d\n", person.GetName(), person.GetAge())
}