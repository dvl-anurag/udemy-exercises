package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)
	fmt.Printf("x and y are %v and %v\t", x, y)

	switch {
	case x < 4 && y < 4:
		fmt.Println("both are less and four")
	case x > 6 || y > 6:
		fmt.Println("both are greater then six")
	case x >= 4 && y <= 6:
		fmt.Println("x is from 4 to 6 inclisive")
	default:
		fmt.Println("None of the previous were met")
	}
}
