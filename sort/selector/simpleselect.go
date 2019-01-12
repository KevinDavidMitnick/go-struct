package selector

func SimpleSelect(l []int) []int {
	length := len(l)
	for i := 0; i < length; i++ {
		min := i
		for j := i; j < length; j++ {
			if l[j] < l[min] {
				min = j
			}
		}
		l[i], l[min] = l[min], l[i]
	}

	return l
}
