package grok

import "testing"

func TestCountdown(t *testing.T) {
	// TODO: mock fmt.Print* to verify
	countdown(5)
}

func TestGreet(t *testing.T) {
	// TODO: mock fmt.Print* to verify
	greet("adit")
}

var ftests = []struct {
	name string
	in   int
	out  int
}{
	{"handles zero", 0, 1},
	{"handles one", 1, 1},
	{"three factorial", 3, 6},
	{"five factorial", 5, 120},
}

func TestFact(t *testing.T) {
	for _, tt := range ftests {
		t.Run(tt.name, func(t *testing.T) {
			actual := fact(tt.in)

			if actual != tt.out {
				t.Errorf("Expected the factorial of %d to return %d but instead got %d!", tt.in, tt.out, actual)
			}
		})
	}
}
