package src

import "fmt"
func main()  {
	var x int32 = -5
	var y int8 = 58960 //constant 58960 overflows int8 ,
	var s uint16 = -1 //constant -1 overflows uint16
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(s)
}
