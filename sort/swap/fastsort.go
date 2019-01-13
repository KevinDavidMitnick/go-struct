package swap

func FastSort(l []int) []int {
	length := len(l)
	if length <= 1 {
		return l
	}
	base := l[0]
	i, j := 0, length-1
	for i < j {
		for ; l[j] > base && j > i; j-- {
		}
		for ; l[i] <= base && i < j; i++ {
		}
		l[i], l[j] = l[j], l[i]
	}
	l[i], l[0] = l[0], l[i]
	FastSort(l[0:i])
	FastSort(l[i+1:])
	return l
}
