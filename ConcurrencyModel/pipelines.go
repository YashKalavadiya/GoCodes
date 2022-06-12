package main

import "fmt"

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

func sq(data<-chan int) <-chan int {
	finalData := make(chan int)
	go func() {
		for n := range data {
			finalData <- n*n
		}
		close(finalData)
	}()
	return finalData
}

func main() {

	nums := []int{2,3,4,5,6}

	//Stage 1
	dataChannel := sliceToChannel(nums)

	//Stage 2
	finalChannel := sq(dataChannel)

	//Stage 3
	for n := range finalChannel {
		fmt.Println(n)
	}
}