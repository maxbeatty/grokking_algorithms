package grok

import "testing"

func TestCheckVoter(t *testing.T) {
	var ctests = []struct {
		name     string
		person   string
		expected string
	}{
		{"mike can vote", "mike", "let them vote!"},
		{"mike has voted", "mike", "kick them out!"},
	}

	for _, tt := range ctests {
		t.Run(tt.name, func(t *testing.T) {
			actual := checkVoter(tt.person)

			if actual != tt.expected {
				t.Errorf("Expected %s but received %s", tt.expected, actual)
			}
		})
	}
}
