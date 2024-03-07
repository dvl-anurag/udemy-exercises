package main

import "fmt"

func main() {
	am := map[string]int{
		"Todd":   42,
		"Henry":  16,
		"Padget": 14,
	}

	fmt.Println("The age of Henry was", am["Henry"])

	fmt.Println(am)
	fmt.Printf("%#v\n", am)

	an := make(map[string]int)
	an["Lucas"] = 28
	an["Steph"] = 37
	an["George"] = 78
	fmt.Println(an)
	fmt.Printf("%#v\n", an)
	fmt.Println(len(an))

	// for range over a MAP
	for k, v := range an {
		fmt.Println(k, v)
	}

	for _, v := range an {
		fmt.Println(v)
	}

	for k := range an {
		fmt.Println(k)
	}

	// for range with a SLICE
	xy := []int{42, 43, 44}

	for i, v := range xy {
		fmt.Println(i, v)
	}

	for _, v := range xy {
		fmt.Println(v)
	}

	for i := range xy {
		fmt.Println(i)
	}

	delete(an, "George")
	fmt.Println(an)
	delete(an, "George")

	fmt.Println(an["George"])

	// v, ok := an["Lucas"]
	// if ok {
	// 	fmt.Println("the value prints",v)
	// } else {
	// 	fmt.Println("Key didn't exist")
	// }

	if v, ok := an["Georgey"]; !ok {
		fmt.Println("Key didn't exist")
	} else {
		fmt.Println("the value prints", v)
	}

}
