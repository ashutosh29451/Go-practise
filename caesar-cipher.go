package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter sentence to encrypt")
	text, _ := reader.ReadString('\n')
	runes := []rune(text)
	for i, _ := range runes {
		if runes[i] != ' ' {
			runes[i] = rune(int(runes[i]) + 3)
		}
	}
	s := string(runes)
	fmt.Println(s)
}
