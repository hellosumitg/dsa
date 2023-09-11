package main

import "fmt"

func main() {
	rows := 5
	cols := 10

	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			if i == 1 || i == rows || j == 1 || j == cols {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
