package swap

func BubbleBigSort(l []int) []int {
	length := len(l)
	for i := 1; i < length; i++ {
		flag := true
		for j := 0; j < length-i && flag; j++ {
			if l[j] > l[j+1] {
				l[j], l[j+1] = l[j+1], l[j]
				flag = false
			}
		}
	}

	return l
}

func BubbleSmallSort(l []int) []int {
	length := len(l)
	for i := 1; i < length; i++ {
		flag := true
		for j := length - 1; j >= i && flag; j-- {
			if l[j] < l[j-1] {
				l[j], l[j-1] = l[j-1], l[j]
				flag = false
			}
		}
	}

	return l
}
