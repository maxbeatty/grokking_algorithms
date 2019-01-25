package grok

import (
	"reflect"
	"testing"
)

func TestKnapsackGreedy(t *testing.T) {
	var ktests = []struct {
		name        string
		items       []item
		weightLimit int
		expected    []string
	}{
		{"9.1", []item{
			{"stereo", 3000, 4},
			{"guitar", 1500, 1},
			{"laptop", 2000, 3},
			{"mp3 player", 1000, 1},
			{"iphone", 2000, 1},
		}, 4, []string{"stereo"}},

		{"9.2", []item{
			{"water", 10, 3},
			{"book", 3, 1},
			{"food", 9, 2},
			{"jacket", 5, 2},
			{"camera", 6, 1},
		}, 6, []string{"water", "food", "camera"}},
	}

	for _, tt := range ktests {
		t.Run(tt.name, func(t *testing.T) {
			actual := knapsackGreedy(tt.items, tt.weightLimit)

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v but received %v", tt.expected, actual)
			}
		})
	}
}

func TestKnapsackDyna(t *testing.T) {
	var ktests = []struct {
		name        string
		items       []item
		weightLimit int
		expected    []string
	}{
		{"9.1", []item{
			{"stereo", 3000, 4},
			{"guitar", 1500, 1},
			{"laptop", 2000, 3},
			{"mp3 player", 1000, 1},
			{"iphone", 2000, 1},
		}, 4, []string{"iphone", "mp3 player", "guitar"}},

		{"9.2", []item{
			{"water", 10, 3},
			{"book", 3, 1},
			{"food", 9, 2},
			{"jacket", 5, 2},
			{"camera", 6, 1},
		}, 6, []string{"camera", "food", "water"}},
	}

	for _, tt := range ktests {
		t.Run(tt.name, func(t *testing.T) {
			actual := knapsackDyna(tt.items, tt.weightLimit)

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expected %v but received %v", tt.expected, actual)
			}
		})
	}

}
