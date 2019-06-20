package search

import (
	"fmt"
)

// CircleSearch test
// for example 6 7 8 9 1  2 3 4 5
func CircleSearch(arr []int, find int, low int, high int) int {
	for low <= high {
		mid := (low + high) / 2
		fmt.Println("--------mid is:", mid)
		if arr[mid] == find {
			return mid
		}
		if mid == 0 || mid == high {
			return -1
		}
		if arr[mid] < find {
			low = mid + 1
		}
		if arr[mid] > find {

			if arr[mid+1] > find {
				high = mid - 1
			} else {
				if arr[high] < find {
					high = mid - 1
				} else {
					low = mid + 1
				}

			}
		}
	}

	return -1
}
