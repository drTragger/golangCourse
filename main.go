package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	word := "Привет Вася"
	firstChar := string([]rune(word)[0])
	lastChar := string([]rune(word)[utf8.RuneCountInString(word)-1:])
	fmt.Printf("First char: %s\nLast char: %s\n", firstChar, lastChar)
}
