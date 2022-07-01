package main

import "fmt"

// Practice some code examples

func main() {
	fmt.Println("Hello world")
	fmt.Println(Sum(5))
}

func Sum(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	return sum
}
