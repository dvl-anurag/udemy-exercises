package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("Loded>>>")
}

func main() {
	x := rand.Intn(400)
	fmt.Printf("The value of x is %v\t", x)

	switch {
	case x <= 100:
		fmt.Println("less than 100")
	case x >= 100 && x < 200:
		fmt.Println("between 100 and 199 inclusive")
	default:
		fmt.Println("this is default")

	}
}
