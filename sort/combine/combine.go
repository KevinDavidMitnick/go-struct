package combine

// Combine func
func Combine(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	// find mid position
	mid := length / 2
	// children left combine
	left := Combine(arr[0:mid])

	// children right combine
	right := Combine(arr[mid:])

	// join left join and right child
	var all []int
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			all = append(all, left[i])
			i++
		} else {
			all = append(all, right[j])
			j++
		}
	}
	if i == len(left) && j < len(right) {
		all = append(all, right[j:]...)
	}
	if i < len(left) && j == len(right) {
		all = append(all, left[i:]...)
	}

	return all
}
