package fib

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	type test struct {
		value  int
		answer int
	}
	tests := []test{
		test{1, 2},
		test{3, 5},
	}
	for _, v := range tests {
		_answer := Fibonacci(v.value)
		if _answer != v.answer {
			t.Error(`expect`, v.answer, `got`, _answer)
		}
	}
}

func ExampleFibonacci() {
	// first argument is target
	Fibonacci(1) // output = 1
	Fibonacci(3) // output = 5
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(i)
	}
}
