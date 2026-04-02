package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter string of words")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading string")
		log.Fatal(err)
	}
	count := 0
	runes := []rune(text)
	needchar := true
	//needspace := false
	for i := 0; i < len(runes); i++ {
		if needchar {
			if runes[i] == ' ' || runes[i] == '\n' || runes[i] == '\r' {

			} else {
				needchar = false
				//needspace = true
				count++
			}
		} else {
			if runes[i] != ' ' && runes[i] != '\n' && runes[i] == '\r' {

			} else {
				needchar = true
				//needspace = false
			}
		}
	}
	words := strings.Fields(text)
	fmt.Println(len(words))
	fmt.Println(count)
}
