package main

import (
	"fmt"
	"os"
)

func main() {
	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	name := os.Getenv("USER") // Current Logged in user
	course := "CTO 101"       // Course being viewed

	fmt.Println("\nHi", name, "you're currently watching", course)
}
