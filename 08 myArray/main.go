package main

import "fmt"

func main() {
	fmt.Println("Array Example")

	var fruitList [4]string

	fruitList[0] = "Mango"
	fruitList[1] = "JackFruit"
	fruitList[3] = "Peach"
	
	// fruitList[2] = ""

	fmt.Println("Fruit List: ", fruitList)

	var vegList = []string{"Potato", "Beans", "Cabbage"}

	fmt.Println("Vegy List: ", len(vegList))
	vegList = append(vegList, "Mushroom")

	fmt.Println("Vegy List: ", len(vegList))


}