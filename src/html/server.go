package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

type Student struct {
	Name string
	age int
}
func main()  {
	student := Student{"Radhouen" , 26}
	template, err := template.ParseFiles("hello.html")
	if err != nil {
		log.Fatal("Error:", err)
	}
	err1 := template.Execute(os.Stdout, student)
	if err1!= nil {
		log.Fatal("error excute :", err1)
	}
	fmt.Println("end of project")
}
