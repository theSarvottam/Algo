package main

import (
	"algo/recursion"
	"fmt"
)

func main() {

	var num int = 6
	var fact int = recursion.Factorial(num)

	result := recursion.Multiplication(3, 5)

	fib := recursion.Fibonacci(num)

	fmt.Println(fact)
	fmt.Printf("multiplication %v", result)
	fmt.Println()
	fmt.Printf("fibonacci of num %v = %v", num, fib)

}
