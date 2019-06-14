package swap

//BubbleSmallSort , small bubble sort
func BubbleSmallSort(arr []int) []int {
	length := len(arr)
	for s := 1; s < length; s++ {
		flag := false
		for i := 0; i < length-s; i++ {
			if arr[i] > arr[i+1] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return arr
}

//BubbleBigSort , big bubble sort
func BubbleBigSort(arr []int) []int {
	length := len(arr)
	for s := 1; s < length; s++ {
		flag := false
		for i := 0; i < length-s; i++ {
			if arr[i] < arr[i+1] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return arr
}
