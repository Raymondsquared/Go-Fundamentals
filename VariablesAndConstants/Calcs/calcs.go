package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 10.0000000000
	b := 3

	fmt.Println("\nA is a type", reflect.TypeOf(a), "and B is of type", reflect.TypeOf(b))

	c := int(a) + b
	fmt.Println("\nC has value:", c, "and is of type:", reflect.TypeOf(c))

}
