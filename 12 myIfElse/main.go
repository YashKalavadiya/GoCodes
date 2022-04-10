package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	var result string
	loginCount := 19

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 && loginCount < 20 {
		result = "Watch out"
	} else {
		result = "definately watch out"
	}

	if 9%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not < 10")
	}


	fmt.Println(result)



}