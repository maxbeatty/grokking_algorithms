package grok

import "testing"

var bstests = []struct {
	name     string
	list     []int
	item     int
	expected int
}{
	{"finds index", []int{1, 3, 5, 7, 9}, 3, 1},
	{"does not find", []int{1, 3, 5, 7, 9}, 11, -1},
}

func TestBinarySearch(t *testing.T) {
	for _, tt := range bstests {
		t.Run(tt.name, func(t *testing.T) {
			actual := binarySearch(tt.list, tt.item)

			if actual != tt.expected {
				t.Errorf("Expected the binary search of %v to return %d but instead got %d!", tt.list, tt.expected, actual)
			}
		})
	}
}
