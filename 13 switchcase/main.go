package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Switchcase in golang")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1

	fmt.Println("Value of dice is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Value of dice is 1 and you can open")
	case 2:
		fmt.Println("Value of dice is 2")
	case 3:
		fmt.Println("Value of dice is 3")
	case 4:
		fmt.Println("Value of dice is 4")
		fallthrough
	case 5:
		fmt.Println("Value of dice is 5")
	case 6:
		fmt.Println("Value of dice is 6 and you get another chance")
		
	}
}