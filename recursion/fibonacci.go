package recursion

func fibonacci(num int) int {
	if num == 0 || num == 1 {
		return num
	}
	return fibonacci(num-1) + fibonacci(num-2)
}

func Fibonacci(n int) int {
	return fibonacci(n)
}
