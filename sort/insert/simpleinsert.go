package insert

// SimpleInsert ,simple insert func.
func SimpleInsert(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}

	// select one from unsort arr,insert to the right positon,from right.
	for j := 1; j < length; j++ {
		i := j
		hold := arr[j]
		for i > 0 && hold < arr[i-1] {
			arr[i] = arr[i-1]
			i--
		}
		arr[i] = hold
	}
	return arr
}
