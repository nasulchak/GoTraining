package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Deffered statement")
	fmt.Println("Starting the main function")

	// Exit with status code of 1
	os.Exit(1)

	// This will never be executed
	fmt.Println("end of main function")
}
