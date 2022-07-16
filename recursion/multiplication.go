package recursion

func multiplication(num1 int, num2 int) int {
	if num2 == 1 {
		return num1

	}
	return num1 + multiplication(num1, num2-1)
}

func Multiplication(m int, n int) int {
	return multiplication(m, n)
}
