package main

import "fmt"

func main() {
	
	defer fmt.Println("Hello World 1")
	defer fmt.Println("Hello World 2")
	fmt.Println("Hello")


	myDiffer()
	
	reccDefer(10)
}

func myDiffer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func reccDefer(num int) {
	defer fmt.Println(num)
	if(num == 1) {
		return
	}
	reccDefer(num-1)

}