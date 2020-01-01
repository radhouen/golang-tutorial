package main

import "fmt"
import "lab1/src/reverse-string"
import "lab1/src/github.com/advanced-maths"

func main() {

    // Creating and initializing strings using var keyword
	var str1 string
	str1 = "Go"
    var str2 string
	str2 = "-MyCode !!!"
    // Concatenating strings  Using + operator
	fmt.Println( str1+str2)
    // Creating and initializing strings  Using shorthand declaration
	str3 := "Golang"
	str4 := "Course"
	result := str3 + " " + str4
    fmt.Println( result)
	fmt.Println(reverse_string.Reverse("GOLANG Course"))
	fmt.Println("Tangent 50 = ", advanced_maths.Tan(90) )

}
