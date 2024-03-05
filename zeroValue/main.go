package main

import (
	"fmt"
	"math"
)

const Pi = 3.14

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	////////Type conversion

	var x, y int = 3, 4
	var j float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(j)
	fmt.Println(x, y, z, j)

	//////////////////Type interface

	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)

	////////////////////// CONSTANT
	const World = "🎈World!🎈"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
