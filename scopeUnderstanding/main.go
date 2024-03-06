package main

import (
	"fmt"
	"scope-understanding/tempPackage"

	"github.com/dvl-anurag/puppy"
)

var x = 87

func main() {
	fmt.Println(x)

	printMe()

	y := 65
	fmt.Println(y)

	p1 := Person{
		"Sam",
	}
	p1.sayHello()

	x := 32
	fmt.Println(x)

	printMe()

	if z := 82; z > 50 {
		fmt.Println(z)
	}
	fmt.Println("############################")
	tempPackage.FrontMet()
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())

}

func printMe() {
	fmt.Println(x)
}

type Person struct {
	first string
}

func (p Person) sayHello() {
	fmt.Println("Hii, my name is " + p.first)
}
