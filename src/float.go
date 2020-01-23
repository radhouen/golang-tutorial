package src

import (
	"fmt"
)

func main()  {
	fmt.Println("float example")
	var x float64 = 10
	var y float64 = 120
	var z float64 = x + y // invalid operation (mismatched type float32 , float64)
	//to resolve this error replace float32 by float64
	fmt.Println("x + y =",x + y )
	fmt.Println("x - y =",x - y )
	fmt.Println("x * y =",x * y)
	fmt.Println("x / y =",x / y)
	//fmt.Println("x % y =",x % y) // % is not defined for float64
}

