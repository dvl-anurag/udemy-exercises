package main

import "fmt"

func main() {

	fmt.Println("This is first statemement to run ")
	fmt.Println("Thisd is sec")
	x := 40
	y := 45
	fmt.Printf("x=%v \n y=%v \n", x, y)

	if x < 42 {
		fmt.Println("x is less than 42")
	}

	if x < 42 {
		fmt.Println("x is still less than 42")
	} else {
		fmt.Println("x is not less than 42")
	}

	if x < 42 {
		fmt.Println("x is still less than 42")
	} else if x == 42 {
		fmt.Println("x is equal to 42")
	} else {
		fmt.Println("x is greater than 42")
	}
}
