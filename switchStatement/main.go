package main

import "fmt"

func main() {

	x := 42

	switch {
	case x < 10:
		fmt.Println("x is greter then 10\n")
	case x == 42:
		fmt.Printf("x is equal to 42\n")
	default:
		fmt.Println("this is the default case\n")
	}

	switch x {
	case 40:
		fmt.Println("it is 40")
	case 42:
		fmt.Println("it is 42")
	default:
		fmt.Println("it is default")
	}

	switch x {
	case 40:
		fmt.Println("it is 40")
		fallthrough
	case 42:
		fmt.Println("it is 42")
		fallthrough
	default:

		fmt.Println("it is default")
	}
}
