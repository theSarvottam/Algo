package recursion

func gcd(num1 int, num2 int) int {
	if num1 == 0 {
		return num2
	} else if num2 == 0 {
		return num1
	} else {
		return gcd(num2%num1, num1)
	}
}

func GCD(m, n int) int {
	return gcd(m, n)
}
