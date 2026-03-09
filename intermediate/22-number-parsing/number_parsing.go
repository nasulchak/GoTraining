package main

import (
	"fmt"
	"strconv"
)

func main() {

	numStr := "12345"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error patsing the value: ", err)
	}
	fmt.Println("Parsed Integer:", num)

	numistr, err := strconv.ParseInt(numStr, 10, 32)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
	}
	fmt.Println("Parsed integer:", numistr)

	floatstr := "3.14"
	floatval, err := strconv.ParseFloat(floatstr, 64)
	if err != nil {
		fmt.Println("Error parsing value:", err)
	}
	fmt.Printf("Parsed float: %.2f\n", floatval)

	binaryStr := "1010" // 8 + 0 + 2 + 0 = 10
	decimal, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		fmt.Println("Error parsing binary value:", err)
	}
	fmt.Printf("Parsed binary to decimal: %d\n", decimal)

	hexStr := "FF" // 15 * 16^1 + 15 * 16^0
	hex, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		fmt.Println("Error parsing hex value:", err)
	}
	fmt.Printf("Parsed hex to decimal: %d\n", hex)

	invalidNum := "22FF"
	invalidParse, err := strconv.Atoi(invalidNum)
	if err != nil {
		fmt.Println("Error parsing value:", err)
	}
	fmt.Printf("Parsed invalid number: %d\n", invalidParse)

}
