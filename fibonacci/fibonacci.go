package main

import "fmt"


func Fibonacci(places int) (r []int) {
    a := 0
    b := 1
    for i := 0; i < places; i++ {
        r = append(r, a)
        a, b = b, a+b
    }
    return
}


func main() {
    fmt.Println(Fibonacci(10))
}