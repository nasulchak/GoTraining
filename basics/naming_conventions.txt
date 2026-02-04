package main

import "fmt"

type EmployeeGoogle struct {
	FistName string
	LastName string
	Age      int
}

type EmployeeApple struct {
	FistName string
	LastName string
	Age      int
}

func main() {
	// PascalCase
	// Eq, CalculateArea, UserInfo, NewHTTPRequest
	// Structs, interfaces, enums

	// snake_case
	// Eg, user_id, first_name, http_request

	// UPPERCASE
	// Use case is constants

	// mixedCase (cammelCase)
	// Eg. javaScript, htmlDocument, isValid

	const MAXRETRIES = 5

	var employeeID = 1001

	fmt.Println("EmloyeeID: ", employeeID)
}
