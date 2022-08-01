package main

import (
	divideandconquer "algo/divideAndConquer"
	"algo/recursion"
	"fmt"
)

func main() {

	var num1 int = 6
	var num2 int = 5
	var fact int = recursion.Factorial(num1)

	result := recursion.Multiplication(num1, num2)

	fib := recursion.Fibonacci(num2)

	gcd := recursion.GCD(num1, num2)

	inputArray := []int{10, 30, 3, 56, 21}

	max, min := divideandconquer.FindMaxMinUsingDAC(inputArray)

	fmt.Println(max, min)

	fmt.Println(fact)
	fmt.Printf("multiplication %v\n", result)
	fmt.Printf("%vth fibonacci num = %v\n", num2, fib)
	fmt.Printf("GCD of numbers num1 = %v, num2 = %v,  is %v", num1, num2, gcd)

}
