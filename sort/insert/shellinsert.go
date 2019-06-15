package insert

func ShellInsert(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}

	for step := length / 2; step >= 1; step = step / 2 {
		for i := step; i < length; i += step {
			j := i
			hold := arr[i]
			for ; j > 0 && hold < arr[j-step]; j -= step {
				arr[j] = arr[j-step]
			}
			arr[j] = hold
		}

	}
	return arr
}
