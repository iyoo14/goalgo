package main

import (
    "fmt"
    "os"
    "strconv"
    )

func main() {
    n,_ := strconv.Atoi(os.Args[1])
    fmt.Println(fact(n))
}

func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n*fact(n-1)
}
