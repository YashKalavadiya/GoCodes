package main

import "fmt"
func main() {
	fmt.Println("Structs in golang")
	//no inheritance in golang; no super or parent


	yash := User{"Yash","yashkalavadiya1010@gmail.com",true,19}

	fmt.Println(yash)

	fmt.Printf("Yash details are: %+v\n", yash)
}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}