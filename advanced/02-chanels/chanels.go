package main

import (
	"fmt"
)

func main() {

	// variable := make(chan type)
	greeting := make(chan string)
	greetString := "Hello"

	go func() {
		greeting <- greetString // blocking because it is continuosly trying to receive values , if is ready to receive continuois flow of data.
		greeting <- "World"
		// fmt.Println("End of goroutine.")
		for _, e := range "abcde" {
			greeting <- "Alphabet:" + string(e)
		}
	}()

	// go func() {
	// 	receiver := <-greeting
	// 	fmt.Println(receiver)
	// 	receiver = <-greeting
	// 	fmt.Println(receiver)
	// }()

	receiver := <-greeting
	fmt.Println(receiver)
	receiver = <-greeting
	fmt.Println(receiver)

	for range 5 {
		rcvr := <-greeting
		fmt.Println(rcvr)
	}

	// time.Sleep(1 * time.Second)
	fmt.Println("End of program.")
}
