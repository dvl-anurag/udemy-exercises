package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int64  `json:"Age"`
}

func main() {
	jsonData := `[{"First":"John","Last":"Doe","Age":30},{"First":"Johnsingh","Last":"Doe","Age":30}]`
	fmt.Println(jsonData)

	var people []Person

	err := json.Unmarshal([]byte(jsonData), &people)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(people)

	for i := 0; i < len(people); i++ {
		person := people[i]
		fmt.Printf("Name: %s %s\n", person.First, person.Last)
		fmt.Printf("Age: %d\n", person.Age)
		fmt.Println()
	}

}
