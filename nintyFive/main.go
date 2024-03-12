package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First   string
	Last    string
	Age     string
	Sayings []string
}

func main() {
	u1 := user{
		First:   "John",
		Last:    "Doe",
		Age:     "30",
		Sayings: []string{"JohnDoeJohnDoe", "JohnDoe", "JohnDoe"},
	}
	u2 := user{
		First:   "Johnsingh",
		Last:    "Doe",
		Age:     "30",
		Sayings: []string{"JohnDoeJohnDoe", "JohnDoe", "JohnDoe"},
	}

	users := []user{u1, u2}

	fmt.Println(users)

	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("Cannot encode JSON")

	}
	fmt.Println(users)
}
