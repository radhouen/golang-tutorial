package main

import "fmt"

func multipleTwoFive(n int)  {
	for i:=2; i<=n;i++  {
		if i%2==0 && i%5==0 {
			fmt.Println(i)
		}
	}

}

func multipleTwoThreeFive(n int)  {
	for i:=2; i<=n;i++  {
		if i%2==0 && i%5==0 && i%3==0{
			fmt.Println(i)
		}
	}

}

func main()  {
	multipleTwoFive(30)
	fmt.Println("--------------------- Separator between the two app call --------------------")
	multipleTwoThreeFive(30)
}

