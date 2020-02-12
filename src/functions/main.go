package main

import (
	"fmt"
	"math"
)

func test() string {
	return "Hello nesrina"
}

func FahrenheitIntoCelsius(f float32) (c float32) {
	c = (f - 32) * 5 / 9
	return
}
func feetIntoMeters(m float32) (f float32) {
	f = m * 0.3048
	return
}

func newton(x float64, y float64, n float64) (a float64) {
	var k float64
	for k=0; k<=n;k++ {
		a += nkfactorial(k,n) * math.Pow(x,k) * math.Pow(y, n-k)
		//fmt.Println("k = %v , n= %v, nkfactorial=%v", )
	}
	return
}
func nkfactorial(k float64, n float64) (kn float64)  {

	kn = factorial(n) / (factorial(k) * factorial(n-k))
	return

}
func factorial(n float64) (f float64) {
	if n == 0 || n == 1 {
		f = 1
	} else {
		f = n * factorial(n-1)
	}
	return
}
func main() {
	fmt.Println(FahrenheitIntoCelsius(20))
	fmt.Println(feetIntoMeters(20))
	fmt.Println("factorial 5 :", factorial(5))
	fmt.Println("newton 2:", newton(2,2,2))
	fmt.Println("newton 2:", newton(2,3,3))
}
