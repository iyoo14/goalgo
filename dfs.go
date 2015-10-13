package main

import (
    "fmt"
    "strconv"
    "strings"
    "flag"
    )

var ary []int
var n int
var k int

func main() {
    np := flag.Int("n", 4, "help message for n")
    kp := flag.Int("k", 13, "help message for k")
    ap := flag.String("a", "1,2,4,7", "help message for a")
    flag.Parse()

    n = *np
    tary := strings.Split(*ap, ",")
    ary = make([]int, len(tary))
    for i, kv := range tary {
        ikv, _ := strconv.Atoi(kv)
        ary[i] = ikv
    }
    k = *kp
    fmt.Println(n)
    fmt.Println(ary)
    fmt.Println(k)
    if dfs(0,0) {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}

func dfs(i int, sum int) bool {
    fmt.Printf("sum:%d\n", sum)
    if n == i {
        return sum == k
    }
    if dfs(i + 1, sum) {
        return true
    }
    if dfs(i + 1, sum + ary[i]) {
        return true
    }
    return false
}

