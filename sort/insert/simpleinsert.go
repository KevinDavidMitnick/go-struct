package insert

// SimpleInsert ,simple insert func.
func SimpleInsert(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}

	// select one from unsort arr,insert to the right positon,from right.
	for j := 1; j < length; j++ {
		i := j - 1
		hold := arr[j]
		for i >= 0 && arr[j] < arr[i] {
			arr[i+1] = arr[i]
			i--
		}
		arr[i+1] = hold
	}
	return arr
}
