package fib

// import (
// 	"fmt"
// )

func _fibonacci(target int, prev int, next int, count int) int {
	// fmt.Println(target, prev, next, count)
	if prev >= target {
		return count
	}
	_next := prev + next
	return _fibonacci(target, next, _next, count+1)
}

func Fibonacci(target int) int {
	// since default argument is not supported by golang, separated func is a must
	return _fibonacci(target, 0, 1, 1)
}
