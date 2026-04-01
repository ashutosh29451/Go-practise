package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("Enter a string to reverse")
	var s string
	fmt.Scanln(&s)
	l, r := 0, utf8.RuneCountInString(s)-1
	runes := []rune(s)
	for l < r {
		pl, pr := runes[l], runes[r]
		runes[l] = pr
		runes[r] = pl
		l++
		r--
	}
	fmt.Println(string(runes))
}
