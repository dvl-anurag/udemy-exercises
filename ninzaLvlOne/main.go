package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	First string
	Last  string
	Age   int64
}

func main() {

	usr1 := User{
		First: "John",
		Last:  "Doe",
		Age:   30,
	}
	usr2 := User{
		First: "Johnsingh",
		Last:  "Doe",
		Age:   30,
	}

	usr := []User{usr1, usr2}

	js, err := json.Marshal(usr)
	if err != nil {
		fmt.Println("Error marshalling JSON")
		return
	}

	fmt.Println(string(js))

}
