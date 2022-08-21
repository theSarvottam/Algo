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

// A element in given array said to be majority element if it
// appear in array more than n/2 times.
func CheckMajoryExit(inputSlice []int) int {

	lengthOfArray := len(inputSlice)
	mid := lengthOfArray - 1/2

	midElement := inputSlice[mid]

	startIndex := findStartIndex(inputSlice, midElement, 0, mid-1)
	lastIndex := findLastINdex(inputSlice, midElement, mid+1, lengthOfArray-1)

	if (lastIndex - startIndex) > lengthOfArray/2 {
		return midElement
	}
	return -1

}

func findStartIndex(inputSlice []int, midElement, firstIndex, lastIndex int) int {

	startIndex := -1
	mid := firstIndex + lastIndex/2
	if inputSlice[mid] == midElement {
		if inputSlice[mid-1] == midElement {
			startIndex = findStartIndex(inputSlice, midElement, firstIndex, mid-1)
		} else if inputSlice[mid+1] == midElement {
			startIndex = mid
		}

	} else {
		startIndex = findStartIndex(inputSlice, midElement, mid+1, lastIndex)
	}

	return startIndex

}
func findLastINdex(inputSlice []int, midElement, startIndex, endIndex int) int {

	lastIndex := -1

	return lastIndex

}
