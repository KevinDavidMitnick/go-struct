package selector

func SimpleSelect(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	for i := 0; i < length; i++ {
		min := i
		for j := i; j < length; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}
