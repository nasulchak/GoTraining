package main

import "fmt"

func main() {
	fmt.Println("Factorial of 5 =", factorial(5))
	fmt.Println("Factorial of 10 =", factorial(10))

	fmt.Println("Summ of digits 173 =", summOfDigits(173))
	fmt.Println("Summ of digits 888 =", summOfDigits(888))

	cache := make(map[int]int)
	fmt.Println("Fibonacci number of 777 =", fib(777, cache))

}

func factorial(n int) int {
	// Base case: factorial of 0 is 1
	if n == 0 {
		return 1
	}

	//Recursive case: factorial of n is n * factorial(n-1)
	return n * factorial(n-1)
	// n * (n - 1) * (n - 2) * factorial(n - 3).... factorial(0)
}

func summOfDigits(n int) int {
	// Base case
	if n < 10 {
		return n
	}

	return n%10 + summOfDigits(n/10)
}

func fib(n int, cache map[int]int) int {
	// Base case
	if n <= 1 {
		return n
	}

	if value, ok := cache[n]; ok {
		return value
	}

	result := fib(n-1, cache) + fib(n-2, cache)
	cache[n] = result
	return result
}
