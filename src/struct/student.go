package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	firstName string
	lastName string
	age int
}

func main()  {


	pStudent := new(Student)
	pStudent.firstName = "Radhouen"
	pStudent.lastName = "Assakra"
	fmt.Println("pointer to student", pStudent)

	anotherStudent := &Student{"GoMyCode" , "Go course", 12}
	fmt.Println("another pointer", anotherStudent)
	fmt.Println(reflect.TypeOf(anotherStudent))
}
