package main

import "fmt"

func main() {
	fmt.Println("WElcome to pointers ")

	// var ptr *int
	// fmt.Println("value of pointer before initialization is : ", ptr)

	myNummber := 10

	var ptr = &myNummber

	fmt.Println("Address of variable myNumber is: ", ptr, "and value of this variable: ", *ptr)
}