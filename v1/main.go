package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Employee struct {
	Name     string
	Age      int
	Location string
}

func main() {
	// Read the file content into a byte array/slice
	content, err := os.ReadFile("./data.json")

	if err != nil {
		// TODO; handle the error
		fmt.Println("Error")
	}
	//print raw file content and file type
	fmt.Println(string(content))
	fmt.Println(http.DetectContentType(content))

	var emp Employee
	json.Unmarshal(content, &emp)
	fmt.Println("Print Emp Variable")
	fmt.Printf("Employee: %v has %v years old, from %v\n", emp.Name, emp.Age, emp.Location)

}
