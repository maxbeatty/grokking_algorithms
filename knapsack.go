package grok

import (
	"sort"
)

type item struct {
	Name   string
	Value  int
	Weight int
}

func knapsackGreedy(items []item, weightLimit int) (selected []string) {
	sort.Slice(items, func(i, j int) bool {
		return items[i].Value > items[j].Value
	})

	curWeight := 0

	for _, i := range items {
		if curWeight+i.Weight <= weightLimit {
			selected = append(selected, i.Name)
			curWeight += i.Weight
		}
	}

	return selected
}

func knapsackDyna(items []item, weightLimit int) (selected []string) {
	matrix := make([][]int, len(items)+1)

	for i := range matrix {
		matrix[i] = make([]int, weightLimit+1)
	}

	// go through combinations of how we may add items to knapsack
	// define what weight/value we would receive using dynamic programming
	for i := 1; i < len(items); i++ {
		for j := 1; j <= weightLimit; j++ {
			/*
				If the item at the index matching the current row fits within the weight capacity represented by the current column, take the maximum of either:
				The total worth of the items already in the bag or,
				The total worth of all the items in the bag except the item at the previous row index, plus the new item’s worth
			*/

			if items[i-1].Weight <= j {
				valueOne := matrix[i-1][j]
				// worth of this subset without the previous item, and this item instead
				valueTwo := items[i-1].Value + matrix[i-1][j-items[i-1].Weight]

				matrix[i][j] = valueOne
				if valueTwo > valueOne {
					matrix[i][j] = valueTwo
				}
			} else {
				// if the new worth is not more, carry over the previous worth
				matrix[i][j] = matrix[i-1][j]
			}
		}
	}

	i := len(items)
	w := weightLimit

	// trace back the knapsack to see what items we're going to add
	for i > 0 && w > 0 {
		pick := matrix[i][w]

		// we know bottom rows have biggest value
		// if current row (pick) isn't the same value as prev row...
		if pick != matrix[i-1][w] {
			// we want to take the previous item
			selected = append(selected, items[i-1].Name)
			// remaining weight is now...
			w -= items[i-1].Weight

		}
		// keep looking up the table
		i--
	}

	return selected
}
