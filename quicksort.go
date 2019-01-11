package grok

func euclid(w, h int) int {
	short := w
	long := h

	if w > h {
		short = h
		long = w
	}

	if long%short == 0 {
		return short
	}

	// because int rounding down happens automagically
	fit := long / short

	return euclid(long-(short*fit), short)
}

// recursive func
func sum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	if len(arr) == 1 {
		return arr[0]
	}

	return arr[0] + sum(arr[1:])
}

func count(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	return 1 + count(arr[1:])
}

func max(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	next := max(arr[1:])

	if arr[0] > next {
		return arr[0]
	}

	return next
}

func quicksort(arr []int) []int { return []int{} }
