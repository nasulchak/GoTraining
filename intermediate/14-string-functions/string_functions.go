package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "Hello Go Go!"
	fmt.Println(len(str))

	str1 := "Hello"
	str2 := "World"
	result := str1 + " " + str2
	fmt.Println(result)

	fmt.Println(str[0])
	fmt.Println(str[1:7])

	// String Conversion
	num := 18
	str3 := strconv.Itoa(num)
	fmt.Println(len(str3))

	// Strings split
	fruits := "apple, orange, banana"
	fruits1 := "apple-orange-banana"
	parts := strings.Split(fruits, ",")
	parts1 := strings.Split(fruits1, "-")
	fmt.Println(parts)
	fmt.Println(parts1)

	countries := []string{"Germany", "France", "Italy"}
	joined := strings.Join(countries, ", ")
	fmt.Println(joined)

	fmt.Println(strings.Contains(str, "Go"))
	fmt.Println(strings.Contains(str, "Go?"))

	replaced := strings.Replace(str, "Go", "Universe", 1)
	fmt.Println(replaced)

	strwspace := " Hello Everyone! "
	fmt.Println(strwspace)
	fmt.Println(strings.TrimSpace(strwspace))
	fmt.Println(strings.ToUpper(strwspace))
	fmt.Println(strings.ToLower(strwspace))

	fmt.Println(strings.Repeat("foo", 3))

	fmt.Println(strings.Count("Hellow", "H"))
	fmt.Println(strings.Count("Hellow", "l"))
	fmt.Println(strings.HasPrefix("Hello", "He"))
	fmt.Println(strings.HasPrefix("Hello", "Re"))
	fmt.Println(strings.HasSuffix("Hello", "lo"))
	fmt.Println(strings.HasPrefix("Hello", "ol"))

	str5 := "Hello, 123 Go!333"
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(str5, -1)
	fmt.Println(matches)

	str6 := "Hello, らポぽ"
	fmt.Println(utf8.RuneCountInString(str6))

	// STRING BUILDER
	var builder strings.Builder

	// Write some strings
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("world!")

	// Convert builder to a string
	result2 := builder.String()
	fmt.Println(result2)

	// Using Writerune to add a character
	builder.WriteRune(' ')
	builder.WriteString("How are you?")
	result2 = builder.String()
	fmt.Println(result2)

	// Reset the builder
	builder.Reset()
	builder.WriteString("Starting fresh!")
	result2 = builder.String()
	fmt.Println(result2)
}
