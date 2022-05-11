package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {

	// return *os.File = means value in this memory locatoin ?
	jsonContent, err := os.Open("./data.json")
	if err != nil {
		// Print the error if exist
		fmt.Println(err)
	}

	fmt.Println(jsonContent)                 //Output: &{0xc00007a120} !! Is
	fmt.Println(reflect.TypeOf(jsonContent)) // Output: *os.File

	P := 5

	fmt.Println(P)                     // Output: 5
	fmt.Println(reflect.TypeOf(P))     // Output: int
	fmt.Println(&P)                    // Output: 0xc0000140a8
	fmt.Println(reflect.TypeOf(&P))    // Output: *int
	fmt.Println(*(&P))                 // Output: 5
	fmt.Println(reflect.TypeOf(*(&P))) // Output: int
	defer jsonContent.Close()

	byteContent, _ := ioutil.ReadAll(jsonContent)

	// Decoding json
	var users Users
	json.Unmarshal(byteContent, &users)

	// print parsed json content
	for i := 0; i < len(users.Users); i++ {
		fmt.Printf("User type: %v\n is %v years old. Social Media: %v\n", users.Users[i].Type, users.Users[i].Age, users.Users[i].Social)

	}
	// what does it mean the interface type in inputs.
	a, _ := json.Marshal([]int{1, 2, 3})
	fmt.Println(string(a))

}
