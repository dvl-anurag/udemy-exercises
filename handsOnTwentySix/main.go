package main

import "fmt"

func main() {

	i := 20
	for i > 0 {
		fmt.Println(i)
		i--
	}
	//////////// Twenty seven
	x := 20

	for {
		if x < 0 {
			break
		}
		fmt.Println(x)
		x--
	}

	///////////////// Twenty eight

	for k := 0; k < 100; k++ {
		if k%2 != 0 {

			fmt.Println(k)
		}
	}

}
