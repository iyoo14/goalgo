package main

import (
    "fmt"
    "io"
    "os"
    "strconv"
    )

var top int
var S [1000]int
var MAX int = len(S)

func initialize() {
    top = 0
}

func isEmpty() bool {
    return top == 0
}

func isFull() bool {
    return top >= MAX - 1
}

func push(x int) {
    if isFull() {
        err("over flow")
    }
    top++
    S[top] = x
}

func pop() int {
    if isEmpty() {
        err("under flow")
    }
    top--
    return S[top+1]
}

func err(msg string) {
    fmt.Fprintln(os.Stderr, "%s\n", msg)
    os.Exit(1)
}

func main() {
    initialize()
    var n string
    for {
        _, err := fmt.Scanf("%s", &n);
        if err == io.EOF {
            break
        }
        fmt.Println(n)
        switch n {
            case "+":
                a := pop()
                b := pop()
                fmt.Printf("+ : %d + %d\n",a , b) 
                push(a + b)
            case "-":
                a := pop()
                b := pop()
                fmt.Printf("- : %d - %d\n",a , b) 
                push(b - a)
            case "*":
                a := pop()
                b := pop()
                fmt.Printf("* : %d * %d\n",a , b) 
                push(a * b)
            default:
                i, _ := strconv.Atoi(n)
                fmt.Printf("i: %d\n", i)
                push(i)
        }
    }
    fmt.Println(pop())
}
