package main

import "fmt"

func printer(i int) {
	if i%3 == 0 && i%5 == 0 {
		fmt.Println("FizzBuzz")
	} else if i%3 == 0 {
		fmt.Println("Fizz")
	} else if i%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(i)
	}
}
func main() {
	for i := 1; i <= 100; i++ {
		printer(i)
	}
}
