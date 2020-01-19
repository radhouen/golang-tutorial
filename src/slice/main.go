package main

import "fmt"

func main() {

	// 3 slots of int
	var s = make([]int, 3)
	fmt.Println(s) // [0 0 0]

	// 3 slots of int, capacity of 9
	var s2 = make([]int, 3, 9)
	fmt.Println(s2) // [0 0 0]
}
