package grok

func setCovering(stations map[string][]string, statesNeeded []string) []string {
	finalStations := map[string]bool{}

	for len(statesNeeded) > 0 {
		var bestStation string
		var statesCovered []string

		for station, states := range stations {
			covered := intersection(states, statesNeeded)

			if len(covered) > len(statesCovered) {
				bestStation = station
				statesCovered = covered
			}
		}

		statesNeeded = diff(statesNeeded, statesCovered)
		finalStations[bestStation] = true
	}

	keys := make([]string, 0, len(finalStations))
	for k := range finalStations {
		keys = append(keys, k)
	}

	return keys
}

func intersection(aa []string, bb []string) []string {
	var intersect []string

	for _, a := range aa {
		for _, b := range bb {
			if a == b {
				intersect = append(intersect, a)
			}
		}
	}

	return intersect
}

func diff(aa []string, bb []string) []string {
	var diff []string

	for _, b := range bb {
		for i, a := range aa {
			if a == b {
				diff = append(aa[:i], aa[i+1:]...)
			}
		}
	}

	return diff
}
