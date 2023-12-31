package main

import "fmt"

func main() {
	var rows, cols int

	fmt.Print("Enter the number of rows: ")
	fmt.Scan(&rows)

	fmt.Print("Enter the number of columns: ")
	fmt.Scan(&cols)

	// Loop through each row and column
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			// If it's the first or last row or the first or last column, print an asterisk
			if i == 1 || i == rows || j == 1 || j == cols {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
