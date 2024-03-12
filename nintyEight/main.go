package main

import (
	"fmt"
	"sort"
)

func main() {

	a := []int{4, 6, 7, 88, 11, 40}
	s := []string{"k", "f", "c", "O", "e", "b"}

	fmt.Println(a)
	sort.Ints(a)
	fmt.Println(a)

	fmt.Println(s)
	sort.Strings(s)
	fmt.Println(s)

}
