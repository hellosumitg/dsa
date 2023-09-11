package main

import "fmt"

func main() {
    rows := 5
    cols := 10

    for i := 1; i <= rows; i++ {
        for j := 1; j <= cols; j++ {
            fmt.Print("* ")
        }
        fmt.Println()
    }
}
