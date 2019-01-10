package grok

func binarySearch(list []int, i int) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2

		if list[mid] == i {
			return mid
		}

		if list[mid] < i {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1 // Not Found
}
