package src

import "fmt"

func add (a int, b int) int {
	somme := a + b
	return somme
}

func rectProps(length, width float64)(float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

/*func rectProps(length, width float64)(area, perimeter float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return // no explicit return value
	//area and perimeter are the named return values in the above function.
}
*/
func main()  {
	somme := add(3,5)
	fmt.Println("somme =", somme)

	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f", area, perimeter)
}

