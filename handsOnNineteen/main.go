package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)
	fmt.Printf("the value of x is %v\t", x)

	if x <= 100 {
		fmt.Println("less then 100")
	} else if x >= 101 && x <= 200 {
		fmt.Println("between 101 and 200")
	} else {
		fmt.Println("this was more then 250")
	}
}
