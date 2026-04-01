package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter a string")
	var s string
	fmt.Scanln(&s)
	runes := []rune(s)
	l, r := 0, len(runes)-1
	var ispallindrome bool = true
	for l < r {
		if runes[l] != runes[r] {
			ispallindrome = false
			break
		}
		l++
		r--
	}
	if ispallindrome {
		fmt.Println("Pallindrome")
	} else {
		fmt.Println("Not Pallindrome")
	}
}
