package main

import (
    "fmt"
    )

func main() {
    r := [...]int{6,3,1,5,4,}
    fmt.Println(r)
    N := len(r)
    i := 0
    flag := true 
    for flag {
        flag = false
        for j := N-1; j > i; j-- {
            if r[j] < r[j-1] {
                tmp := r[j]
                r[j] = r[j-1]
                r[j-1] = tmp
                flag = true
            }
        }
        i++
    }
    fmt.Println(r)
}

