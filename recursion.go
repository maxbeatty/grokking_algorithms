package grok

import "fmt"

// countdown using recursion
func countdown(i int) {}

// stack function calls
func greet2(name string) {
	fmt.Printf("how are you, %s?\n", name)
}

func bye() {
	fmt.Println("ok bye!")
}

func greet(name string) {
	fmt.Printf("hello %s!\n", name)

	greet2(name)

	fmt.Println("getting ready to say bye...")

	bye()
}

// factorial using recursion
func fact(i int) int {
	if i == 0 || i == 1 {
		return 1
	}

	return i * fact(i-1)
}
