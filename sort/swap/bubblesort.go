package swap

func BubbleBigSort(l []int) []int {
	length := len(l)
	for i := 1; i < length; i++ {
		flag := true
		for j := 0; j < length-i; j++ {
			if l[j] > l[j+1] {
				l[j], l[j+1] = l[j+1], l[j]
				flag = false
			}
		}
		if flag == true {
			break
		}
	}

	return l
}

func BubbleSmallSort(l []int) []int {
	length := len(l)
	for i := 1; i < length; i++ {
		flag := true
		for j := length - 1; j >= i; j-- {
			if l[j] < l[j-1] {
				l[j], l[j-1] = l[j-1], l[j]
				flag = false
			}
		}
		if flag == true {
			break
		}
	}

	return l
}
