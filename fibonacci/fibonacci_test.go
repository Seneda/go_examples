package main

import "testing"

var result []int

func runBenchmark(b *testing.B, fib func(int) []int, places int) {
	var r []int
	for i:= 0; i < b.N; i++ {
		r = fib(places)
	}
	result = r
}

// Benchmarks for testing how much faster it is to preallocate a slice of the
// correct size vs appending to the slice for each new place calculated

func BenchmarkFibonacciAppShort(b *testing.B) {
	runBenchmark(b, FibonacciApp, 10)
}

func BenchmarkFibonacciAllShort(b *testing.B) {
	runBenchmark(b, FibonacciAll, 10)
}

func BenchmarkFibonacciAppLong(b *testing.B) {
	runBenchmark(b, FibonacciApp, 1e4)
}

func BenchmarkFibonacciAllLong(b *testing.B) {
	runBenchmark(b, FibonacciAll, 1e4)
}

func BenchmarkFibonacciAppVeryLong(b *testing.B) {
	runBenchmark(b, FibonacciApp, 1e7)
}

func BenchmarkFibonacciAllVeryLong(b *testing.B) {
	runBenchmark(b, FibonacciAll, 1e7)
}
