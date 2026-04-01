package main

import (
	"fmt"
)

func fibo(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fibo(n-1) + fibo(n-2)
}
func main() {
	var a int
	fmt.Println("Enter which fibonacci number you want")
	fmt.Scan(&a)
	fmt.Println(fibo(a))
}
