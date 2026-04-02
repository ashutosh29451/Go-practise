package main

import "fmt"

func main() {
	fmt.Println("Number of elements")
	n := 0
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	if n > 0 {
		ans := arr[0]
		for _, num := range arr {
			ans = max(ans, num)
		}
		fmt.Println(ans)
	}
}
