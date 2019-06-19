package search

// Search1 ,for test.
func Search1(arr []int, element int, low int, high int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if element == arr[mid] {
		return mid
	}
	if element < arr[mid] {
		return Search1(arr, element, low, mid-1)
	}
	if element > arr[mid] {
		return Search1(arr, element, mid+1, high)
	}
	return -1
}

// Search2 , for test.
func Search2(arr []int, element int, low int, high int) int {
	for low < high {
		mid := (low + high) / 2
		if element == arr[mid] {
			return mid
		}
		if element < arr[mid] {
			high = mid - 1
		}
		if element > arr[mid] {
			low = mid + 1
		}

	}

	return -1
}
