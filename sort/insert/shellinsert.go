package insert

func ShellInsert(l []int) []int {
	length := len(l)

	for step := length / 2; step >= 1; step /= 2 {
		for i := step; i < length; i += step {
			elem := l[i]
			var j = i
			for ; j > 0 && elem < l[j-step]; j -= step {
				l[j] = l[j-step]
			}
			l[j] = elem
		}
	}

	return l
}
