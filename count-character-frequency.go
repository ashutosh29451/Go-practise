package main

import "fmt"

func main() {
	s := ""
	fmt.Println("enter string")
	fmt.Scanln(&s)
	mymap := make(map[rune]int)
	for _, c := range s {
		mymap[c]++
	}
	for c, f := range mymap {
		fmt.Printf("%c : %d \n", c, f)
	}
}
