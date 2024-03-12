package starting_code

import (
	"fmt"
	//"github.com/dvl-anurag/cat"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		//age:  cat.Years(10),
	}
	fmt.Println(fido)
	//fmt.Println(cat.YearsTwo(20))
}
