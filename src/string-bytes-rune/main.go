package main

import "fmt"

func main() {

	// print a rune in decimal, hex, and standard unicode notations

	var capA = 'A'

	fmt.Printf("%d %x %U\n", capA, capA, capA) // 65 41 U+0041
}
