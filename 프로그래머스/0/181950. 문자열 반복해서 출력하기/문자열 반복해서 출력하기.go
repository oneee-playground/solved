package main

import (
    "fmt"
    "strconv"
)

func main() {
    var s1 string
    var a string
    fmt.Scan(&s1, &a)
    
    n, _ := strconv.Atoi(a)
    for i := 0; i < n; i++ {
        fmt.Printf(s1)
    }
    fmt.Println()
}