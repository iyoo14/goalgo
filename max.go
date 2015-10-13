package main

import (
    "fmt"
    )

func main() {
    var maxv int
    var minv int
    r := [...]int{6,5,3,1,3,4,3}
    l := len(r)
    minv = r[0]
    for j := 1; j < l; j++  {
        minv = min(minv, r[j])
        maxv = max(maxv, r[j] - minv)
    }
    fmt.Printf("Max : %d\n", maxv)
}

func max(i int, j int) (int) {
    if i > j {
        return i
    }
    return j
}

func min(i int, j int) (int) {
    if i < j {
        return i
    }
    return j
}
