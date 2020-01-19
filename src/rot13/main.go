package main

import (
	"fmt"
	"strings"
)

func rot13(r rune) rune {
	fmt.Println("rune element :", r)
	if r >= 'a' && r <= 'z' {
		if r >= 'm' {
			return r - 13
		} else {
			return r + 13
		}
	} else if r >= 'A' && r <= 'Z' {
		if r >= 'M' {
			return r - 13
		} else {
			return r + 13
		}
	}
	return r
}
func main() {
	input := "Do you have any questions ?"
	mapped := strings.Map(rot13, input)
	fmt.Println(input)
	fmt.Println(mapped)
}
