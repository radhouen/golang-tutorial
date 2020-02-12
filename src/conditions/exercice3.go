package main

import "fmt"

func prime(n int)   {
	for i:=2;i<n;i++ {
		p := true
		for j:=2;j<i;j++ {
			if i%j == 0 {
				p = false
			}
		}
		if p == true {
			fmt.Println("primary : %d", i)
		}
	}
}

func main()  {
	prime(100)
}
