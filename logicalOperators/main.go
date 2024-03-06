package main

import "fmt"

func main() {
	x := 40
	y := 100
	fmt.Printf(" x=%v \n y=%v\n", x, y)

	if x < 42 && y < 42 {
		fmt.Println("both are less then 42")
	}

	if x > 30 || x < 42 {
		fmt.Println("x is getting close to 42")
	}

	if x != 42 && y != 10 {
		fmt.Println("x is not 42 and y is not 10")
	}
}
