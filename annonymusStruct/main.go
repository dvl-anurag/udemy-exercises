package main

import "fmt"

type Person struct {
	first string
	last  string
	age   int
}

func main() {
	sa1 := struct {
		first string
		last  string
		age   int
	}{
		first: "Jenny",
		last:  "Moneypenny",
		age:   27,
	}
	p2 := Person{
		first: "Jenny",
		last:  "Moneypenny",
		age:   27,
	}

	fmt.Println(sa1)
	fmt.Println(p2)

	fmt.Printf("%T \t %#v\n", sa1, sa1)

	fmt.Println(sa1.first, sa1.last, sa1.age)
	fmt.Println(sa1.first, sa1.last, sa1.age)
	fmt.Println(sa1.first)
}
