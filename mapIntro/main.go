package main

import "fmt"

func main() {
	am := map[string]int{
		"Todd":   42,
		"Henry":  16,
		"padget": 14,
	}

	fmt.Println("The age of Henry was", am["Henry"])

	fmt.Println(am)
	fmt.Printf("%v\n", am)

	an := make(map[string]int)
	an["Lucas"] = 42
	an["Steph"] = 16
	fmt.Println(an)

	fmt.Println(len(an))

}
