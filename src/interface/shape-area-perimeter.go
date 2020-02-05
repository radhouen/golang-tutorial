package main

import (
	"math"
	"fmt"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rect) Area() float64 {
	return  r.width * r.height
}
func (r Circle) Area() float64 {
	return math.Pi * r.radius * r.radius

}
func (r Rect) Perimeter() float64 {
	return  r.width + r.height
}
func (r Circle) Perimeter() float64 {
	return 2 * math.Pi * r.radius
}

func main()  {
	var r Shape = Rect{10,3}

	fmt.Printf("type of r is %T\n:", r)
	fmt.Printf("value of r is %v\n:", r)
	fmt.Printf("Area of r is %v\n:", r.Area())
	fmt.Printf("Perimeter of r is %v\n:", r.Perimeter())


	var c  = Circle{5}
	fmt.Printf("type of c is %T\n:", c)
	fmt.Printf("value  of c is %v\n:", c)
	fmt.Printf("Area of c is %v\n:", c.Area())
	fmt.Printf("Perimeter of c is %v\n:", c.Perimeter())

}
