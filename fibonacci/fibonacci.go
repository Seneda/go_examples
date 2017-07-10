package main

import (
    "fmt"
)

func FibonacciApp(places int) (r []int) {
    // Append each new term onto the old slice
    a := 0
    b := 1
    for i := 0; i < places; i++ {
        r = append(r, a)
        a, b = b, a+b
    }
    return
}

func FibonacciAll(places int) (r []int) {
    // Preallocate a slice of the full size required
    a := 0
    b := 1
    r = make([]int, places)
    for i := 0; i < places; i++ {
        r[i] = a
        a, b = b, a+b
    }
    return
}

func Fibonacci(places int) []int {
    return FibonacciAll(places)
}


func main() {
    fmt.Println(Fibonacci(100))
}