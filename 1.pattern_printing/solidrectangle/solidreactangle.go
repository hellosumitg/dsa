package main

import "fmt"

func main() {
	var rows, cols int

	fmt.Print("Enter the number of rows: ")
	fmt.Scan(&rows)

	fmt.Print("Enter the number of columns: ")
	fmt.Scan(&cols)

	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
