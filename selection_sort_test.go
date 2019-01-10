package grok

import (
	"reflect"
	"testing"
)

var sstests = []struct {
	name     string
	list     []int
	expected []int
}{
	{"sorts asc", []int{5, 3, 10, 2, 6}, []int{2, 3, 5, 6, 10}},
}

func TestSelectionSort(t *testing.T) {
	for _, tt := range sstests {
		t.Run(tt.name, func(t *testing.T) {
			actual := selectionSort(tt.list)

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v but instead received %v", tt.expected, actual)
			}
		})
	}
}
