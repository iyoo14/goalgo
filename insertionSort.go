package main

import (
    "fmt"
    )

func main() {
    r := [...]int{6,5,3,1,3,4,3}
    fmt.Println(r)
    l := len(r)
    for i := 1; i < l; i++  {
        v := r[i]
        j := i -1
        for  (j >= 0 &&  v < r[j]) {
            r[j+1] = r[j]
            j--
        }
    fmt.Println(j)
        r[j+1] = v
    }
    fmt.Println(r)
}

