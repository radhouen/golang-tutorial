package main

import (
	"fmt"
	"lab1/src/github.com/advanced-maths"
	"lab1/src/reverse-string"
	"strings"
)

func main() {
	fmt.Println("replace character : %v", replaceCharacter("e", "w", "Radhouen"))
	fmt.Println("replace character : %v", replaceCharacter("French", "India", "BlaBlaCar is a French online marketplace for carpooling"))
	fmt.Println("replace vowel by 1  function result : %v", replaceVowel("radhouen"))
	fmt.Println("Delete special characters : %v =", deleteSpecialCharacter("radhouen%-Assakra."))
	fmt.Println("Delete special characters : %v =", deleteSpecialCharacter("You can find functions or install plugins to generate Lorem Ipsum text in your text editing software like Microsoft Word, OpenOffice, Notepad ++ or Lotus Notes; your presentation software like Keynote or Pages; your CMS like Joomla, Drupa, Mambo, PH-Nuke, WordPress or Movable Type; your web browsers like Chrome; your desgin software like Photosh"))
	// Creating and initializing strings using var keyword
	var str1 string
	str1 = "Go"
	var str2 string
	str2 = "-MyCode !!!"
	// Concatenating strings  Using + operator
	fmt.Println(str1 + str2)
	// Creating and initializing strings  Using shorthand declaration
	str3 := "Golang"
	str4 := "Course"
	result := str3 + " " + str4
	fmt.Println(result)
	fmt.Println(reverse_string.Reverse("GOLANG Course"))
	fmt.Println("Tangent 50 = ", advanced_maths.Tan(90))
	containsString("Seafood", "food")
	printSpecialCharacter()

}
func Hello() (s string) {
	s = "Hello World !"
	return
}

func upperCase(s string) {

}
func compareString(s string, s1 string) {

}
func containsString(s string, s1 string) bool {
	fmt.Println("Contains function", strings.Contains(s, s1))
	return strings.Contains(s, s1)
}
func equal(s string, s1 string) bool {
	return strings.EqualFold(s, s1)
}
func replace(s string, s1 string) {

}
func replaceAll(s string, s1 string) {

}
func split(s string, s1 string) {

}
func toLowers(s string, s1 string) {

}

func length(s string) {

}
func join(str [] string, sep string) string {
	return strings.Join(str, sep)
}

func mapStrings(mapping func(rune) rune, str string) string {
	mapping = func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	fmt.Println(strings.Map(mapping, str))
	return strings.Map(mapping, str)
}
func printSpecialCharacter() {
	var x = "♥\t😂"
	fmt.Printf("%s\n", x)  // ♥ 😂
	fmt.Printf("%q\n", x)  // "♥\t😂"
	fmt.Printf("%+q\n", x) // "\u2665\t\U0001f602"
	fmt.Printf("% x\n", x) // e2 99 a5 09 f0 9f 98 82
}

func replaceVowel(s string) (str string) {
	return strings.Map(func(r rune) rune {
		switch r {
		case 'a':
			return '1'
		case 'e':
			return '1'
		case 'i':
			return '1'
		case 'u':
			return '1'
		case 'o':
			return '1'
		case 'y':
			return '1'

		}
		return r
	}, s)
}

func replaceCharacter(oldStr string, newStr string, s string) (str string) {
	i := strings.Index(s, oldStr)
	fmt.Println("first occurance : %d", i)
	str = strings.Replace(s, oldStr, newStr, i)
	return

}

func deleteSpecialCharacter(s string) (str string) {
	str = strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' {
			return r
		} else {
			return rune(1)
		}
	}, s)
	return str
}
