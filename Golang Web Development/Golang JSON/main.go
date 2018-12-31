package main 

import (
	"fmt"
	"encoding/json"
)
type person struct {
	First	string `json:"First"`
	Last	string	`json:"Last"`	
	Age		int		`Json:"Age"`
}
func main() {
	s := `[{"First": "srikant", "Last": "Prasad" "age": 32},"{First": "Sai", "Last": "Prasad", "age": 32}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	person := []person{}
	// var person []person
	err := json.Unmarshal(bs, &person)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("all of the data", person)
	for i, v := range person {
		fmt.Println("\n Person number", i)
		fmt.Println(v.First, v.Last, v.Age)
	}

	fmt.Println("hello world!")
}