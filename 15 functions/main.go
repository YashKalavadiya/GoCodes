package main

import "fmt"

func main() {
	greeter()

	func () {
		fmt.Println("IFFE")
	}()

	result := adder(2,5);
	fmt.Println("Result is: ", result)

	allAdditions, msg := proAdder(2,3,4,5,6,7,8)
	fmt.Println("All additions are: ", allAdditions, msg)
	
	
}

func greeter() {
	fmt.Println("Hello from yash")
}

func adder(n1 int, n2 int) int {
	return n1+n2
}

func proAdder (values ...int) (int, string) {
	total := 0

	for _,val := range values {
		total += val
	}
	return total, "Pro result"
}