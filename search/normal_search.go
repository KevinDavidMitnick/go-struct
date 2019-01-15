package search

func NormalSearch(arr []int, element int) int {
	length := len(arr)
	i := 0
	for ; i < length; i++ {
		if element == arr[i] {
			break
		}
	}
	if i < length {
		return i
	}
	return -1
}
