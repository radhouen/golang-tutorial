package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	firstName string
	lastName  string
	age       int // default value  = 0
}

var Students = [] Student{
	{firstName: "Radhouen", lastName: "Assakra", age: 26},
	{firstName: "Radhouen", lastName: "Assakra", age: 26},
	{firstName: "Radhouen", lastName: "Assakra", age: 26},
}

func main() {

	pStudent := new(Student)
	pStudent.firstName = "Radhouen"
	pStudent.lastName = "Assakra"
	fmt.Println("pointer to student", pStudent)

	anotherStudent := &Student{"GoMyCode", "Go course", 12}
	fmt.Println("another pointer", anotherStudent)
	fmt.Println(reflect.TypeOf(anotherStudent))
	for _, student := range Students {
		fmt.Println("user name :", student.firstName)
		fmt.Println("user name :", student.lastName)
		fmt.Println("user name :", student.age)

	}
}
