package main

import "fmt"

func main() {
	name := "Raymond"   // Name of subscriber
	course := "CTO 101" // Course being viewed

	fmt.Println("\nHi", name, "you're currently watching", course)

	changeCourseByValue(course)
	fmt.Println("\nHi", name, "you're now still watching", course)

	changeCourseByReference(&course)
	fmt.Println("\nHi", name, "you're now watching", course)
}

func changeCourseByValue(course string) string {
	course = "COO 101"

	fmt.Println("\nTrying to change your course by value to", course)

	return course
}

func changeCourseByReference(course *string) string {
	*course = "CFO 101"

	fmt.Println("\nTrying to change your course by reference to", *course)

	return *course
}
