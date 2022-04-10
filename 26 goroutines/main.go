package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup  //usually these are pointers

var signals = []string {"test"}

var mut sync.Mutex //usually pointer


func main() {
	// go greeter("Hello")
	// greeter("world")

	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()

	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i:= 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string){
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOp!!! ")
	}else {
		mut.Lock() //lock memory so that no other routine can write to it
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("\n%d status code for %s", res.StatusCode, endpoint)
	}

}

//Go routines can access same memory for reading puspose. but not for writing same memory (if we use mutex)
