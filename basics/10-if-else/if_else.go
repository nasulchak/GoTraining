package main

import "fmt"

func main() {
	age := 25

	if age >= 18 {
		fmt.Println("You are an adult.")
	}

	// if contidion {
	// block of code
	//} else if {
	// block of code
	//} else {
	// block of code
	//}

	temperature := 25

	if temperature >= 30 {
		fmt.Println("It`s hot outside")
	} else {
		fmt.Println("It`s cool outside")
	}

	score := 85

	if score >= 90 {
		fmt.Println("Grade A")
	} else if score >= 80 {
		fmt.Println("Grade B")
	} else if score >= 70 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Grade D")
	}

	// if contidion1 {
	//	 Code of block
	//	 if contidion2 {
	//		Code of block2
	//	}
	//}

	num := 18
	if num%2 == 0 {
		if num%3 == 0 {
			fmt.Println("Number is divisible by both 2 and 3.")
		} else {
			fmt.Println("Number is divisible by 2 bot not 3.")
		}
	} else {
		fmt.Println("Numer is not divisible by 2.")
	}

	// || OR
	// && AND

	if 10%2 == 0 || 5%2 == 0 {
		fmt.Println("Either 10 or 5 are even.")
	}

	if 10%2 == 0 && 6%2 == 0 {
		fmt.Println("Both 10 or 6 are even.")
	}
}
