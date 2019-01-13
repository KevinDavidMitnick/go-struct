package insert

func SimpleInsert(l []int) []int {
	length := len(l)

	for i := 1; i < length; i++ {
		elem := l[i]
		var j = i
		for ; j > 0 && elem < l[j-1]; j-- {
			l[j] = l[j-1]
		}
		l[j] = elem
	}

	return l
}
