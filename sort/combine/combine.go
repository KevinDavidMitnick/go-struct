package combine

func Combine(l []int) []int {
	length := len(l)
	if length <= 1 {
		return l
	}

	mid := length / 2
	left := Combine(l[0:mid])
	right := Combine(l[mid:])
	return mergeSort(left, right)
}

func mergeSort(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}
