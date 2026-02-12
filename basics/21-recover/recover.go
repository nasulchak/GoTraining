package main

import "fmt"

func main() {
	process()
	fmt.Println("Returned from pocess")
}

func process() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	fmt.Println("Start Process")
	panic("Somethind went wrong!")
}
