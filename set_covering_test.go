package grok

import (
	"testing"
)

func TestSetCovering(t *testing.T) {
	statesNeeded := []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"}

	stations := map[string][]string{
		"kone":   []string{"id", "nv", "ut"},
		"ktwo":   []string{"wa", "id", "mt"},
		"kthree": []string{"or", "nv", "ca"},
		"kfour":  []string{"nv", "ut"},
		"kfive":  []string{"ca", "az"},
	}

	actual := setCovering(stations, statesNeeded)

	// contents of actual shifts around with each run... :(
	if len(actual) != 4 {
		t.Errorf("Expected 4 stations but received %d: %v", len(actual), actual)
	}

}
