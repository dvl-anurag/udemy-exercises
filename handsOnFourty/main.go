package main

import "fmt"

func main() {
	str := [...]string{"hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello"}
	fmt.Println(str)
	fmt.Println(len(str))

	//////////// Slice

	sl := []string{"hello", "hello"}
	fmt.Println(sl)

	////////handsOn fourtyOne

	xs := []string{"hello", "hello", "hello", "hello", "hello"}
	fmt.Println(xs)
	fmt.Printf("%T\n, xs")

	for i, v := range xs {
		fmt.Printf("%v - %v\n", i, v)
	}

	for _, v := range xs {
		fmt.Printf("%v\n", v)

	}

	fmt.Println("-----------")
	fmt.Println(xs[0])
	fmt.Println(xs[1])
	fmt.Println(xs[2])
	fmt.Println(xs[3])

	fmt.Println(len(xs))

	///////////////append to a slice

	xk := []int{1, 2, 3, 4, 5}

	fmt.Println(xk)

	fmt.Println("--------------------------")

	xk = append(xk, 6, 4, 5, 6)

	fmt.Println(xk)

	//////////////////// slicing a slice

	sk := []int{1, 2, 3, 4, 5}

	fmt.Printf("sk - %#v\n", sk)
	fmt.Println("---------------")
	fmt.Printf("sk- %#v\n", sk[0:4])
	fmt.Println("---------------------")
	fmt.Printf("sk - %#v\n", sk[:3])
	fmt.Println("---------------------")
	fmt.Printf("sk - %#v\n", sk)

	fmt.Println(sk[0])

	///////////////Delete from a slice

	sk = append(sk[:2], sk[3:]...)

	fmt.Printf("sk - %#v\n", sk)
	fmt.Println("--------------")

	///////// slice-make

	si := []string{"d", "f", "c"}
	fmt.Println(si)

	zx := make([]int, 0, 10)
	fmt.Println(zx)
	fmt.Println(len(zx))
	fmt.Println(cap(zx))
	zx = append(zx, 5, 4, 5, 3, 6, 4, 6)
	fmt.Println(zx)
}
