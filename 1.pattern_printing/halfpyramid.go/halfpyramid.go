package main

import "fmt"

func main() {
	rows := 5 // Number of rows in the half pyramid

	for i := 1; i <= rows; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
