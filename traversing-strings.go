package main

import "fmt"

func main() {
	str := "Cafés are cool!"
	for i := 0; i < len(str); i++ {
		fmt.Println(i, str[i])
	}
	for i, r := range str {
		fmt.Println(i, r)
	}
}
