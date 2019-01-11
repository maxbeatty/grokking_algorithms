package grok

var voted map[string]bool

func init() {
	voted = make(map[string]bool)
}

func checkVoter(name string) string {
	if voted[name] {
		return "kick them out!"
	}
	voted[name] = true
	return "let them vote!"
}
