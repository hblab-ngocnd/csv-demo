package bench

import (
	"testing"
)

func FibRecursive(n int) int {
	if n < 2 {
		return n
	}
	return FibRecursive(n-1) + FibRecursive(n-2)
}

func BenchmarkFibRecursive10(b *testing.B) {
	// run the FibRecursive function b.N times
	for n := 0; n < b.N; n++ {
		FibRecursive(10)
	}
}

func FibNonRecursive(n int) int {
	if n < 2 {
		return n
	}
	f1 := 0
	f2 := 1
	tmp := f2
	for i := 2; i < n; i++ {
		tmp = f2
		f2 = f1 + f2
		f1 = tmp
	}
	return f2
}

func BenchmarkFibNonRecursive10(b *testing.B) {
	// run the FibNonRecursive function b.N times
	for n := 0; n < b.N; n++ {
		FibNonRecursive(10)
	}
}

func FibNonRecursiveOptimized(n int) int {
	if n < 2 {
		return n
	}
	f1 := 0
	f2 := 1
	tmp := f2
	for i := 2; i < n; i++ {
		tmp = f2
		f2 = f1 + f2
		f1 = tmp
	}
	return f2
}

func BenchmarkFibNonRecursiveOptimized10(b *testing.B) {
	// run the FibNonRecursive function b.N times
	for n := 0; n < b.N; n++ {
		FibNonRecursive(10)
	}
}

func BenchmarkFib10(b *testing.B) {
	// run the FibRecursive function b.N times
	for n := 0; n < b.N; n++ {
		FibRecursive(10)
	}
	for n := 0; n < b.N; n++ {
		FibNonRecursive(10)
	}
	for n := 0; n < b.N; n++ {
		FibNonRecursiveOptimized(10)
	}
}
