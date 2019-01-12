package insert

func SimpleInsert(l []int) []int {
	length := len(l)

	for i := 1; i < length; i++ {
		elem := l[i]
		index := locale(l, elem)
		for j := i; j > index; j-- {
			l[j] = l[j-1]
		}
		l[index] = elem
	}

	return l
}

func locale(l []int, ele int) int {
	length := len(l)
	index := length - 1
	for i := 0; i < length; i++ {
		if ele <= l[i] {
			index = i
			break
		}
	}
	return index
}
