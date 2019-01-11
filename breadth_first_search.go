package grok

func search(graph map[string][]string, name string, check func(string) bool) bool {
	var searchQueue []string
	searchQueue = append(searchQueue, graph[name]...)
	searched := make(map[string]bool)
	var person string

	for len(searchQueue) != 0 {
		person = searchQueue[0]
		searchQueue = searchQueue[1:]

		if !searched[person] {
			if check(person) {
				return true
			}

			searchQueue = append(searchQueue, graph[person]...)
			searched[person] = true
		}
	}
	return false
}
