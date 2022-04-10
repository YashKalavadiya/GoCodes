package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
)

func main() {
	fmt.Println("Welcome to math section")

	// var myNumberOne int = 2
	// var myNumberTwo float64 = 4

	// fmt.Println("The sum is: ", myNumberOne + int(myNumberTwo))

	//RANDOM number

	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5))

	//RANDOM from crypto

	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))

	fmt.Println(myRandomNum)
}