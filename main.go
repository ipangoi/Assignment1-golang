package main

import (
	"assignment1/model"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {

	arg := os.Args[1]
	jsonCall(arg)

}

func jsonCall(input string) {
	// Get the absolute path to the JSON file
	jsonPath := filepath.Join("data", "participants.json")

	// Read JSON file content
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		log.Fatal(err)
	}

	// var person model.Person
	var people model.Participant
	var cond bool

	// Unmarshal JSON data into person struct
	if err := json.Unmarshal(data, &people); err != nil {
		log.Fatal(err)
	}

	for _, person := range people.Participant {
		if person.Code == input {
			cond = true
		}

	}

	// return &model.Participant{}, false

	switch cond {
	case true:
		for _, person := range people.Participant {
			if person.Code == input {
				fmt.Printf("ID: %d\n", person.ID)
				fmt.Printf("Code: %s\n", person.Code)
				fmt.Printf("Name: %s\n", person.Name)
				fmt.Printf("Address: %s\n", person.Address)
				fmt.Printf("Occupation: %s\n", person.Occupation)
				fmt.Printf("Reason Joining: %s\n", person.Reason)
				fmt.Printf("==================\n")
			}
		}
	default:
		fmt.Println("Not Found")
	}
}
