package main

import (
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

	fmt.Println(fact)
	fmt.Printf("multiplication %v\n", result)
	fmt.Printf("%vth fibonacci num = %v\n", num2, fib)
	fmt.Printf("GCD of numbers num1 = %v, num2 = %v,  is %v", num1, num2, gcd)

}
