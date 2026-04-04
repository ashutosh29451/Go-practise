package main

import "fmt"

func main() {
	n, m := 0, 0
	fmt.Println("Number of rows")
	fmt.Scan(&n)
	fmt.Println("number of columns")
	fmt.Scan(&m)
	slice := make([][]int, n)
	for i := 0; i < n; i++ {
		slice[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("Enter element for row %d and column %d \n", i, j)
			fmt.Scan(&slice[i][j])
		}
	}
	flatSlice := make([]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			flatSlice = append(flatSlice, slice[i][j])
		}
	}
	fmt.Println("flat slice is")
	fmt.Println(flatSlice)
}
