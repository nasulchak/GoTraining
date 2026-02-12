package main

import "fmt"

func main() {
	// panic(interface{})

	// Example of a valid input
	// process(10)

	// Exmaple of an invalid input
	process(-10)
}

func process(input int) {
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")

	if input < 0 {
		fmt.Println("Before panic")
		panic("input must be a non-negative number")
	}
	fmt.Println("Processing input:", input)
}
