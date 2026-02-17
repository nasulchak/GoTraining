package main

import (
	"fmt"
	"slices"
	"strings"
	"unicode/utf8"
)

func main() {

	message := "Hello,\nGo!"
	message1 := "Hello,\tGo!"
	message2 := "Hello,\rGo!"
	rawMessage := `Hello\nGo`

	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)

	fmt.Println("Length of message variable is", len(message))
	fmt.Println("Length of message1 variable is", len(message1))
	fmt.Println("Length of message2 variable is", len(message2))
	fmt.Println("Length of rawMessage variable is", len(rawMessage))

	fmt.Println("The first character in message var is", message[0]) // ASCII

	greeting := "Hello "
	name := "Alice"
	fmt.Println(greeting + name)

	str1 := "Apple"  // A has an ASCII value of 65
	str := "apple"   // A has an ASCII value of 65
	str2 := "banana" // b has an ASCII value of 98
	str3 := "app"    // a has an ASCII value 97
	fmt.Println(str1 < str2)
	fmt.Println(str3 < str1)
	fmt.Println(str > str1)
	fmt.Println(str > str3)

	for _, char := range message {
		// fmt.Printf("Character at index %d is %c\n", i, char)
		fmt.Printf("%x-%c\n", char, char)
	}

	fmt.Println("Rune count:", utf8.RuneCountInString(greeting))

	greetingWithName := greeting + name
	fmt.Println(greetingWithName)

	var ch rune = 'a'
	fmt.Println(ch)
	fmt.Printf("%c - %T\n", ch, ch)

	cstr := string(ch)
	fmt.Println(cstr)
	fmt.Printf("Type of cstr is %T\n", cstr)

	r := '😁'
	fmt.Printf("%v\n", r)
	fmt.Printf("%c\n", r)

	s := "hêllo"
	runes := []rune(s)
	fmt.Printf("%c\n", runes[2]) // Третья руна
	for i := range s {
		fmt.Printf("position %d: %c\n", i, s[i])
	}

	fmt.Println(strings.TrimRight("123oxo", "xo"))  // Вернёт "123"
	fmt.Println(strings.TrimSuffix("123oxo", "xo")) // Теперь вернёт "123o"

	fmt.Println(concat([]string{"abc", "123", "def"}))

	fmt.Println(reverse("hello"))  // "olleh"
	fmt.Println(reverse("Привет")) // "тевирП"
	fmt.Println(reverse("🙂🙃"))     // "🙃🙂"

	fmt.Println(countChars("hello"))  // 5
	fmt.Println(countChars("Привет")) // 6
	fmt.Println(countChars("🙂🙃"))     // 2

	fmt.Println(isPalindrome("level")) // true
	fmt.Println(isPalindrome("Level")) // true
	fmt.Println(isPalindrome("топот")) // true
	fmt.Println(isPalindrome("🙂🙃🙂"))   // true
	fmt.Println(isPalindrome("hello")) // false
}

func concat(arr []string) string {
	var sb strings.Builder

	for _, v := range arr {
		sb.WriteString(v)
	}

	return sb.String()
}

func reverse(s string) string {
	runeS := []rune(s)
	slices.Reverse(runeS)
	return string(runeS)
}

func countChars(s string) int {
	return utf8.RuneCountInString(s)
}

func isPalindrome(s string) bool {
	r := []rune(strings.ToLower(s))

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		if r[i] != r[j] {
			return false
		}
	}
	return true
}
