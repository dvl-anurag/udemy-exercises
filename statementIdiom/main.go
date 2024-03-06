package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("first")
	fmt.Println("sec")

	x := 40
	y := 5
	fmt.Printf("x=%v \n y=%v \n", x, y)

	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v \n and that is %T\n", z, x)
	} else {
		fmt.Printf("z is %v and that is %T\n", z, x)
	}
}
