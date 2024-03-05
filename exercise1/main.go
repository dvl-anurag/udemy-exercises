package main

import "fmt"

var x int
var y string
var z bool

var a = 18
var b = "03"
var c = false

func main() {

	fmt.Println("Hello World!😊😊😎🎈")
	fmt.Println(`Udemy India LLP
	10th Floor, ResCowork 07, Tower B,
	Unitech Cyber Park, Sector 39,
	Gurgaon, Haryana, India, 122003
	GSTIN: 06AAFFU9763M1ZE
	PAN no. AAFFU9763M`)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	s := fmt.Sprintf("%v %v %v", a, b, c)
	fmt.Println(s)
}
