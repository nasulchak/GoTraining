package main

import "fmt"

func main() {
	sequence := adder()

	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())

	sequence2 := adder()

	fmt.Println(sequence2())

	subtraction := func() func(int) int {
		countDown := 99
		return func(x int) int {
			countDown -= x
			return countDown
		}
	}()

	fmt.Println(subtraction(10))
	fmt.Println(subtraction(9))
}

func adder() func() int {
	i := 0
	fmt.Println("Previous value of i:", i)

	return func() int {
		i++
		fmt.Println("Added 1 to i")
		return i
	}
}
