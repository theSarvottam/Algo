package divideandconquer

func LinearSearch(inputSlice []int, number int) int {

	for i, data := range inputSlice {
		if data == number {
			return i
		}
	}
	return -1
}

func BinarySearch(inputSlice []int, number int) int {

	index := binarySearch(inputSlice, 0, len(inputSlice)-1, number)
	return index

}

func binarySearch(inputSlice []int, startingIdex int, LastIndex int, number int) int {

	numberIndex := -1
	//check small problem or not
	if startingIdex == LastIndex {
		if inputSlice[startingIdex] == number {
			return startingIdex
		} else {
			return numberIndex
		}
	} else {
		mid := (startingIdex + LastIndex) / 2
		if inputSlice[mid] == number {
			return mid
		} else if inputSlice[mid] > number {
			numberIndex = binarySearch(inputSlice, startingIdex, mid-1, number)
		} else {
			numberIndex = binarySearch(inputSlice, mid+1, LastIndex, number)
		}
	}
	return numberIndex

}


