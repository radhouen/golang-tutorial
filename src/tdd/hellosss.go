package main

import "fmt"
import "lab1/src/tdd/math"

func Hello() string {
	return "Hello, world!"
}
func add(i int, i2 int) int {
	return i - i2
}
func main() {
	fmt.Println(Hello())
	numbers := [5]int{1, 2, 2, 4, 5}
	math.Sum(numbers)
}
