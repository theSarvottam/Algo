package divideandconquer

func StrightMaxMin(inputArray []int) (int, int) {

	max, min := inputArray[0], inputArray[0]
	// for element := range inputArray {
	// 	if element > max {
	// 		max = element
	// 	} else if element < min {
	// 		min = element
	// 	}
	// }
	for i := 1; i < len(inputArray); i++ {
		if inputArray[i] > max {
			max = inputArray[i]
		} else if inputArray[i] < min {
			min = inputArray[i]
		}
	}

	return max, min

}

func FindMaxMinUsingDAC(inputArray []int) (int, int) {
	if smallProblem(inputArray) {
		return solution(inputArray)
	} else {
		var p int = len(inputArray) / 2
		max1, min1 := FindMaxMinUsingDAC(inputArray[:p])
		max2, min2 := FindMaxMinUsingDAC(inputArray[p:])
		return combine(max1, min1, max2, min2)

	}

}

func smallProblem(inputArray []int) bool {
	if len(inputArray) == 1 || len(inputArray) == 2 {
		return true
	}
	return false
}

func solution(inputArray []int) (int, int) {
	if len(inputArray) == 1 {
		return inputArray[0], inputArray[0]
	}
	if inputArray[0] > inputArray[1] {
		return inputArray[0], inputArray[1]
	}
	return inputArray[1], inputArray[0]
}

func combine(max1 int, min1 int, max2 int, min2 int) (max int, min int) {
	if max1 > max2 {
		max = max1
	} else {
		max = max2
	}
	if min1 < min2 {
		min = min1
	} else {
		min = min2
	}
	return max, min
}
