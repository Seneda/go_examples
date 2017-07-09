package main

import "fmt"
import "math"


func Pi() (p float64){
    for i := 1.0; i < 100000000; i = i+4 {
        p += 4 / i
        p -= 4 / (i + 2)
    }
    return
}

func main() {
    p := Pi()
    fmt.Printf("Pi()    : %.15f\n", p)
    fmt.Printf("math.Pi : %.15f\n", math.Pi)
    fmt.Printf("Difference between math.Pi and Pi() is : %e", p-math.Pi)
}