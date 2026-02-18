package main

import "fmt"

func main() {

	//Printing Functions
	fmt.Print("Hello")
	fmt.Print("World!")
	fmt.Print(12, 456)

	fmt.Println("Hello")
	fmt.Println("World!")
	fmt.Println(12, 456)

	name0 := "John"
	age0 := 25
	fmt.Printf("Name: %s, Age: %d\n", name0, age0)
	fmt.Printf("Binary age: %b, Hex age: %X\n", age0, age0)

	// Foramtting Functions
	s := fmt.Sprint("Hello", "World!", 123, 456)
	fmt.Println(s)

	s = fmt.Sprintln("Hello", "World!", 123, 456)
	fmt.Print(s)

	sf := fmt.Sprintf("Name: %s, Age: %d", name0, age0)
	fmt.Println(sf)

	// Scanning Functions
	var name string
	var age int

	fmt.Print("Enter your name and age:")
	fmt.Scan(&name, &age)
	// fmt.Scanln(&name, &age)
	// fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("Name: %s, age: %d\n", name, age)

	// Error Formatting Functions
	err := checkAge(7)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("Age %d is too young to drive.", age)
	}

	return nil
}
