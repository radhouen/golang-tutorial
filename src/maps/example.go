package main

import "fmt"

func main()  {
	 monthdays := map[string]int{
		"Jan": 31, "Feb": 28,"Mar":31,
		"Apr": 30, "May": 31,"Jun":30,
		"Jul": 31, "Aug": 31,"Sep":30,
		"Oct": 31, "Nov": 30,"Dec":31,
	}
	fmt.Println(monthdays)
	fmt.Println("January", monthdays["Jan"])

	year := 0
	for _, days := range monthdays {
		year += days
	}
	fmt.Println("year days = ", year)

}
