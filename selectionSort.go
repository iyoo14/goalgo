package main

import "fmt"


func main() {
    r := [...]int{6,3,1,5,4,}
    fmt.Println(r)
    l := len(r)
    for i := 0; i < l; i++  {
        minj := i
        for j := i+1; j < l; j++ {
            if r[j] < r[minj] {
                minj = j
            }
            tmp := r[i]
            r[i] = r[minj]
            r[minj] = tmp
        }
    }
    fmt.Println(r)
}

