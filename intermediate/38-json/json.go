package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string  `json:"name"`
	Age       int     `json:"age,omitempty"`
	Email     string  `json:"email,omitempty"`
	Address   Address `json:"addres"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func main() {
	person := Person{FirstName: "John"}

	// Marshalling
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	fmt.Println(string(jsonData))

	person1 := Person{FirstName: "Jane", Age: 18, Email: "jane@gmail.com", Address: Address{City: "New York", State: "NY"}}

	jsonData1, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	fmt.Println(string(jsonData1))

	jsonData2 := `{"first_name":"Damir", "age": 19, "address":{"city":"New City", "state":"NC"}}`

	var employeeFromJson Employee
	err = json.Unmarshal([]byte(jsonData2), &employeeFromJson)
	if err != nil {
		fmt.Println("Error unmarshal JSON:", err)
		return
	}

	fmt.Println(employeeFromJson)

	listOfCityState := []Address{
		{City: "New York", State: "NY"},
		{City: "San Jose", State: "CA"},
		{City: "Las Vegas", State: "NV"},
		{City: "Modesto", State: "CA"},
		{City: "Clearwater", State: "FL"},
	}

	fmt.Println(listOfCityState)
	jsonList, err := json.Marshal(listOfCityState)
	if err != nil {
		log.Fatalln("Error Marshaling to JSON:", err)
	}

	fmt.Println("JSON List:", string(jsonList))

	// Handling unknow json structures
	jsonData3 := `{"name": "John", "age": 30, "address": {"city": "New York", "state": "NY"}}`

	var data map[string]interface{}

	err = json.Unmarshal([]byte(jsonData3), &data)
	if err != nil {
		log.Fatalln("Error Marshaling to JSON:", err)
	}
	fmt.Println("Decode/Unmarshalled JSON:", data)
}

type Employee struct {
	FirstName string  `json:"first_name"`
	Age       int     `json:"age"`
	Address   Address `json:"address"`
}
