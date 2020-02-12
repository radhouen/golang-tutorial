package main

import "fmt"

func main() {

	k := 1
	for ; k <= 10; k++ {
		fmt.Println(k)
	}

	k = 1
	for k <= 10 {
		fmt.Println(k)
		k++
	}

	for k := 1; ; k++ {
		fmt.Println(k)
		if k == 10 {
			break
		}
	}
	var a int = 10

	/* do loop execution */
	for a < 20 {
		if a == 15 {
			/* skip the iteration */
			a = a + 1
			//continue
		}
		fmt.Printf("value of a: %d\n", a)
		a++
	}
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("Continuing loop")
			continue // break here
		}
		fmt.Println("The value of i is", i)
	}

}
