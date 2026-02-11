package main

import "fmt"

func main() {
	// ... Ellipsis
	//func functionName(param1 type1, param2 type2, param3 ...type3) returnType {
	// code block
	//}

	sequence, total := sum(1, 1, 2, 3, 4)
	fmt.Println("Sequence: ", sequence, "Total:", total)

	numbers := []int{1, 2, 3, 4, 5, 9}
	sequence2, total2 := sum(2, numbers...)
	fmt.Println("Sequence: ", sequence2, "Total:", total2)

}

func sum(sequence int, nums ...int) (int, int) {
	total := 0
	for _, v := range nums {
		total += v
	}

	return sequence, total
}
