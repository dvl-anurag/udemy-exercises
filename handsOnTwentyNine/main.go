package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println("outerloop %v \t innerloop %v\n", i, j)
		}
		fmt.Println("******************")
	}

	/////////////// Hand on 30
	xi := []int{42, 43, 44, 45, 46, 47}

	for i, v := range xi {
		fmt.Printf("index %v \t value %v\n", i, v)
	}

	/////////////Hands on thirty
	fmt.Println("------------------------------")
	m := map[string]int{
		"James": 43,
		"manni": 32,
	}

	for k, vv := range m {
		fmt.Printf("key %v\t value %v\n", k, vv)
	}

	////////////////////////////////HandsOn thirty one

	age1 := m["James"]
	fmt.Println("Age of James is ", age1)
	if y, ok := m["James"]; !ok {
		fmt.Println("There is no key James, and here is the zero value of an int", y)

	}

	age2 := m["q"]
	fmt.Println("Age of q is ", age2) // will

	if y, ok := m["Q"]; !ok {
		fmt.Println("There is no key Q, and here is the zero value of an int", y)

	}

	//////////////////////HandsOn  Thirty two
	c := 1
	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Printf("iteration %v \t total count %v \t x is %v\n", i, c, x)

		}
	}
}
