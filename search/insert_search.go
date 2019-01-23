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
		return BinarySearch(arr, element, low, high-mid)
	}
	if element > arr[mid] {
		return BinarySearch(arr, element, low+mid, high)
	}
	return -1
}
