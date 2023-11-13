package main

import "fmt"

func main() {
	var rows int

	fmt.Print("Enter the number of rows: ")
	fmt.Scan(&rows)

	// Outer loop to iterate through each row in reverse order (from bottom to top)
	for i := rows; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
	fmt.Println()
	}
}
