package main

import(
    "fmt"
    "os"
    "strconv"
    )

func main() {
    n := os.Args[1]
    in, _ := strconv.Atoi(n)
    fmt.Println(in)
}
