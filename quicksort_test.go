package grok

import "testing"

func TestEuclid(t *testing.T) {
	actual := euclid(1680, 640)

	if actual != 80 {
		t.Errorf("euclid calc expected 80 but received %d", actual)
	}
}

func TestSum(t *testing.T) {
	var stests = []struct {
		name     string
		list     []int
		expected int
	}{
		{"works", []int{2, 4, 6}, 12},
	}

	for _, tt := range stests {
		t.Run(tt.name, func(t *testing.T) {
			actual := sum(tt.list)

			if actual != tt.expected {
				t.Errorf("Expected %d but received %d when summing %v", tt.expected, actual, tt.list)
			}
		})
	}
}

func TestCount(t *testing.T) {
	var ctests = []struct {
		name     string
		list     []int
		expected int
	}{
		{"works", []int{1, 1, 1}, 3},
	}

	for _, tt := range ctests {
		t.Run(tt.name, func(t *testing.T) {
			actual := count(tt.list)

			if actual != tt.expected {
				t.Errorf("Expected %d but received %d when counting %v", tt.expected, actual, tt.list)
			}
		})
	}
}

func TestMax(t *testing.T) {
	var mtests = []struct {
		name     string
		list     []int
		expected int
	}{
		{"works", []int{2, 1, 3}, 3},
		{"more", []int{1, 5, 10, 25, 16, 1}, 25},
	}

	for _, tt := range mtests {
		t.Run(tt.name, func(t *testing.T) {
			actual := max(tt.list)

			if actual != tt.expected {
				t.Errorf("Expected %d but received %d when finding max of %v", tt.expected, actual, tt.list)
			}
		})
	}
}
