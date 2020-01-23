package main

import (
	"fmt"
	"regexp"
	"strings"
)
var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
func main()  {
fmt.Println("result",replaceAllVowell("radhouen"))

	fmt.Println("result",removeSpecialCharacter("r!!!!adhouen"))
}

func replaceAllVowell(str string) string  {

	fmt.Println("str before modification :", str)
	str = strings.Map(func(r rune) rune {
		switch r {
		case 'a':
			return '1'
		case 'e':
			return '1'
		// etc.
		default:
			return r
		}
	}, str)
	return str
}

func removeSpecialCharacter(str string) string  {
	fmt.Println("str before modification :", str)
	str = strings.Map(func(r rune) rune {
		if((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z' )) {
			return r
		} else {
			return rune(0)
		}
	}, str)
	return str
}
