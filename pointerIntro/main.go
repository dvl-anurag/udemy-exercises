package main

import "fmt"

func main() {
	x := 5
	fmt.Println(x)
	fmt.Println(&x)

	fmt.Printf("%v \t %T\n", &x, &x)

	s := "hello"

	fmt.Printf("%v \t %T\n", &s, &s)

	m := 39

	fmt.Printf("%v \t %T\n", &m, &m)

	n := &m

	fmt.Printf("%v \t %T\n", &n, &n)

	fmt.Println(*n)

	fmt.Println(*&m)

	*n = 45

	fmt.Println(m)

	fmt.Println(*n)

}
