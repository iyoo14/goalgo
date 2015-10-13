package main

import (
    "fmt"
    "os"
    "strconv"
    )

const (
    MAX int =  40
    )
var MEMO[MAX+1]int
func main() {
    n,_ := strconv.Atoi(os.Args[1])
    fmt.Println(fib(n))
}

func fib(n int) int {
    if n <= 1 {
        return n
    }
    if MEMO[n] != 0 {
        return MEMO[n]
    }
    MEMO[n] = fib(n-1) + fib(n-2)
    return MEMO[n]
}

