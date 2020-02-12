package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Fprintf(os.Stdout, "%d + %d = %d \n", a, b, sum(a,b))
	fmt.Fprintf(os.Stdout, "%d - %d = %d \n", a, b, minus(a,b))
	fmt.Fprintf(os.Stdout, "%d * %d = %d \n", a, b, mult(a,b))
	fmt.Fprintf(os.Stdout, "%d / %d = %d \n", a, b, div(a,b))
}

func sum(a int, b int) int {
	return a + b
}
func minus(a int, b int) int {
	return a - b
}
func mult(a int, b int) int {
	return a * b
}
func div(a int, b int) int {
	if (b != 0) {
		return a / b
	}
	return 0
}
