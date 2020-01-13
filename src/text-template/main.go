package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type Student struct {
	//exported field since it begins
	//with a capital letter
	Name string
}

func generateTemplate( str string) string  {
	s := Student{Name:str}
	template := template.New("index")
	template, err := template.Parse("Hello {{.Name}}")
	if err != nil {
		fmt.Println(err)
		log.Fatal("Parse", err)
		return ""
	}
	err1 := template.Execute(os.Stdout,s)
	if err1 != nil {
		log.Fatal("excute", err1)
		return ""
	}
	return ""
}

func generateTextTemplate(str string) string  {

	return ""
}

func main() {
	//define an instance
	/*s := Student{"Satish"}

	//create a new template with some name
	tmpl := template.New("test")

	//parse some content and generate a template
	tmpl, err := tmpl.Parse("Hello {{.Name}}!")
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	//merge template 'tmpl' with content of 's'
	err1 := tmpl.Execute(os.Stdout, s)
	if err1 != nil {
		log.Fatal("Execute: ", err1)
		return
	}*/
	generateTemplate("Radhouen")
}
