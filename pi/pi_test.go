package main

import (
    "testing"
    "math"
    "fmt"
)

func TestPiPrecise(t *testing.T) {
    e := math.Abs(math.Pi-Pi())
    lim := 1e-11
    if e > lim {
        t.Error(fmt.Sprintf("Error was %e which is greater than %e", e, lim))
    }
}

func TestPiLessPrecise(t *testing.T) {
     e := math.Abs(math.Pi-Pi())
     lim := 1e-7
     if e > lim {
         t.Error(fmt.Sprintf("Error was %e which is greater than %e", e, lim))
     }
 }

