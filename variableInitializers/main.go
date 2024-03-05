package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, true, "ok"
	fmt.Println(i, j, c, python, java)

	//////////////Short var declaration

	var n, m int = 1, 2
	l := 3
	d, ds, java := true, false, "no!"

	fmt.Println(n, m, l, ds, d, java)
}
