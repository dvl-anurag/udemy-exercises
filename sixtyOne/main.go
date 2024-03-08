package main

import "fmt"

type Person struct {
	first string
	age   int
}

func (p Person) Speak() {
	fmt.Println("My name is", p.first, "and my age is", p.age)
}

func main() {
	p1 := Person{
		first: "Jenny",
		age:   27,
	}

	p1.Speak()

}
