package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in go")

	myCh := make(chan int, 2)//2 means buffered channel, even if multiple values are sent through channel, it will only take 1

	wg := &sync.WaitGroup{}


	// fmt.Println(<-myCh)
	// myCh <- 5

	wg.Add(2)

	//receive ONLY
	go func(ch <-chan int, wg *sync.WaitGroup){

		// close(myCh)
		val, isChanelOpen := <- myCh
		// myCh <- 10
		fmt.Println(isChanelOpen)
		fmt.Println(val)
		fmt.Println(<- myCh)
		defer wg.Done()
	}(myCh, wg)

	//send ONLY
	go func(ch chan<- int, wg *sync.WaitGroup){
		myCh <- 5
		myCh <- 6
		// fmt.Println(<-myCh)
		// close(myCh)
		defer wg.Done()
	}(myCh, wg)

	wg.Wait()
}