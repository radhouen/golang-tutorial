package main

import "fmt"

func calculateSalaries(tab []int) (som float32 )  {

	for _, s := range tab  {
		som += float32(s ) * 1.3
	}
	return
}

func calculateTicketValue( size int, value int, number int) (pt float32)  {

	pt = float32(size * value * number )
	return
}

func main()  {
	tab := []int {1500,1200,1800,2630,1590,1600,1550,1850,1400,1350,1250,1250,2400,1900}
	fmt.Println("Total Amount =", calculateSalaries(tab) + calculateTicketValue(len(tab),10,20))
}
