package main

import "fmt"

func main() {
	var rows int
	fmt.Print("Enter the number of rows: ")
	fmt.Scan(&rows)

	for i := rows; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			// If it's the first or last row or the first or last column, print an asterisk
			if j == 1 || j == i || i == rows {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
