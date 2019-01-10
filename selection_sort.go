package grok

func findSmallestIdx(list []int) int {
	low := list[0]
	lowIdx := 0

	for i := 1; i < len(list); i++ {
		if list[i] < low {
			low = list[i]
			lowIdx = i
		}
	}

	return lowIdx
}

func selectionSort(list []int) []int {
	size := len(list)
	sorted := make([]int, size)

	for i := 0; i < size; i++ {
		smallest := findSmallestIdx(list)
		sorted[i] = list[smallest]
		// new list with appended slices up to but excluding smallest and then everything _after_ smallest
		list = append(list[:smallest], list[smallest+1:]...)
	}

	return sorted
}
