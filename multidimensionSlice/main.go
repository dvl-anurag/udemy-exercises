package main

import "fmt"

func main() {

	jb := []string{"sam", "tam", "game"}
	jm := []string{"kam", "l0am", "lame"}

	fmt.Println(jb)
	fmt.Println(jm)

	fm := [][]string{jb, jm}
	fmt.Println(fm)

	//// slice internals and underlying array

	a := []string{"kam", "l0am", "lame"}

	b := a
	fmt.Println("a", a)
	fmt.Println("b", b)

	a[0] = "9"

	fmt.Println("a", a)
	fmt.Println("b", b)

	//////////slice internals
	c := []string{"kam", "l0am", "lame"}

	d := make([]string, 6)
	copy(d, c)
	fmt.Println("c", c)
	fmt.Println("d", d)

	c[0] = "9"

	fmt.Println("c", c)
	fmt.Println("d", d)

}
