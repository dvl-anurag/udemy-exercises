package main

import "fmt"

func main() {
	y := 12
	for i := 0; i < 5; i++ {
		fmt.Printf("counting to five %v \t first fot loop\n", i)
	}

	for y < 10 {
		fmt.Println("y is %v \t\t\t second for loop\n", y)
	}

	for {
		fmt.Println("y is %v \t\t third for loop\n", y)
		if y > 20 {
			break
		}
		y++
	}

	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("counting even no: ", i)
	}
}
