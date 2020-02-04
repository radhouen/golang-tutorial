package main

import "fmt"

func sum( tab [] int) (som int)  {
	for _, s := range tab {
		som += s
	}
	return
}

func getAverage(arr []int) (avg float32) {

	avg = float32(sum(arr) / len(arr))
	return
}

func main() {
	tab := []int {1,2,3,4,5,6}
	fmt.Println("Somme =", sum(tab) )
	fmt.Println("Average =", getAverage(tab) )
}

