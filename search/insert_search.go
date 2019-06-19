package search

func InsertSearch(arr []int, element int, low int, high int) int {
	length := len(arr)
	if length == 0 {
		return -1
	}
	mid := low + (high-low)*(element-arr[low])/(arr[high]-arr[low])
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
