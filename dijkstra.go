package grok

import (
	"math"
)

// DRY tests
type conn map[string]int

// dag == directed acyclic graph
type dag map[string]conn

func dijkstra(graph dag) (weightShortestPath int) {
	// assumes "fin" is destination node
	processed := map[string]bool{}
	costs := map[string]int{}
	parents := map[string]string{"fin": ""}

	// populate costs and parents
	for node, conns := range graph {
		for con, cos := range conns {
			if _, ok := costs[con]; !ok {
				costs[con] = cos
				parents[con] = node
			}
		}
	}

	costs["fin"] = math.MaxInt8

	findLowestCostNode := func() string {
		lowestCost := math.MaxInt8
		lowestCostNode := ""

		for node, cost := range costs {
			if cost < lowestCost && !processed[node] {
				lowestCost = cost
				lowestCostNode = node
			}
		}

		return lowestCostNode
	}

	var cost int
	var neighbors conn

	node := findLowestCostNode()

	for node != "" {
		cost = costs[node]
		neighbors = graph[node]

		for nname, ncost := range neighbors {
			newCost := cost + ncost

			if costs[nname] > newCost {
				costs[nname] = newCost
				parents[nname] = node
			}

		}

		processed[node] = true
		node = findLowestCostNode()
	}

	return costs["fin"]
}
