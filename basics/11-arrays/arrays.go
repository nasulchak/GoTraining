package main

import "fmt"

func main() {
	// var arrayName [size]elementType

	var numbers [5]int
	fmt.Println(numbers)

	numbers[4] = 20
	fmt.Println(numbers)

	numbers[0] = 9
	fmt.Println(numbers)

	fruits := [4]string{"Apple", "Banana", "Orang", "Grapes"}
	fmt.Println("Fruits array:", fruits)

	fmt.Println("Third element:", fruits[2])

	originalArray := [3]int{1, 2, 3}
	coppiedArray := originalArray

	coppiedArray[0] = 100

	fmt.Println("Original array:", originalArray)
	fmt.Println("Coppied array:", coppiedArray)

	for i := 0; i < len(numbers); i++ {
		fmt.Println("Element at index", i, ":", numbers[i])
	}

	for i, v := range numbers {
		fmt.Printf("Index %d, Value: %d\n", i, v)
	}

	for i := range numbers {
		fmt.Printf("Index: %d\n", i)
	}

	for _, v := range numbers {
		fmt.Printf("Value: %d\n", v)
	}

	fmt.Println("The length of numbers array is", len(numbers))

	// Comparing Arrays
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{10, 2, 3}

	fmt.Println("Array1 is equal to Array2:", array1 == array2)

	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(matrix)

	originArray := [3]int{1, 2, 3}
	var coppyArray *[3]int

	coppyArray = &originArray
	coppyArray[0] = 100
	fmt.Println("OriginArray:", originArray)
	fmt.Println("CoppyArray:", coppyArray)
}
