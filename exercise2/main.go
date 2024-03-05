package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(math.Pi)
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println("The number is", rand.Intn(10))
	fmt.Println(add(92, -13))

}
