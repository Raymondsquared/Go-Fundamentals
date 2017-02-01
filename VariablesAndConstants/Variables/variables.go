package main

import (
	"fmt"
	"reflect"
)

func main() {

	name := "Raymond" // Name of subscriber
	// course := "CTO 101" // Course being viewed
	module := 5.1 // Current position in course
	ptr := &module

	fmt.Println("Name is", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("Module is", module, "and is of type", reflect.TypeOf(module))
	fmt.Println("Memory address if *module* variable is", ptr, "and the value of *module* is", *ptr)
}
