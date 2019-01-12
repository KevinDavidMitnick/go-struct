package insert

func BinaryInsert(l []int) []int {
	length := len(l)

	for i := 1; i < length; i++ {
		elem := l[i]
		index := binaryLocale(l[:i], elem)
		for j := i; j > index; j-- {
			l[j] = l[j-1]
		}
		l[index] = elem
	}

	return l
}

func binaryLocale(l []int, elem int) int {
	length := len(l)
	var low, high int
	for low, high = 0, length-1; low <= high; {
		mid := (low + high) / 2
		if elem < l[mid] {
			high = mid - 1
		}
		if elem > l[mid] {
			low = mid + 1
		}
	}
	return low
}
