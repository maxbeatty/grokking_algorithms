package grok

import "testing"

func TestSearch(t *testing.T) {
	var graph = make(map[string][]string)

	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}

	var check = func(name string) bool {
		return name[len(name)-1] == 'm'
	}

	actual := search(graph, "you", check)

	if !actual {
		t.Errorf("Expected search starting with you to return true")
	}

}
