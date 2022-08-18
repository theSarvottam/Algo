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

	inputArray := []int{10, 30, 40, 56, 80}

	max, min := divideandconquer.FindMaxMinUsingDAC(inputArray)

	elementIndex := divideandconquer.BinarySearch(inputArray, 56)

	fmt.Println(max, min)
	fmt.Println(elementIndex)

	fmt.Println(fact)
	fmt.Printf("multiplication %v\n", result)
	fmt.Printf("%vth fibonacci num = %v\n", num2, fib)
	fmt.Printf("GCD of numbers num1 = %v, num2 = %v,  is %v", num1, num2, gcd)

}
