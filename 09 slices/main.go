package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println("Welcome to slices")

	var fruitList = []string{"Apple", "Tomatoes", "Peach"}

	// fmt.Printf("Type of fruitlist is: %T\n", fruitList)

	// fmt.Println("before append: ", fruitList)
	fruitList = append(fruitList, "Strawberry")
	// fmt.Println("After append: ", fruitList)

	fruitList = append(fruitList[1:])

	// fmt.Println("After append: ", fruitList)


	highScore := make([]int, 4)

	highScore[0] = 222
	highScore[1] = 324
	highScore[2] = 542
	highScore[3] = 678

	highScore = append(highScore, 888, 555, 999)

	// fmt.Println("Prof of append works in adding elements than bound specified: ", highScore)

	sort.Ints(highScore)

	// fmt.Println(highScore)

	//how to remove a value from a slice based on index

	var courses = []string{"React js", "MERN", "javascript", "Python", "Swift"}

	fmt.Println(courses)

	var idx int = 3
	courses = append(courses[:idx], courses[idx+1:]...)
	fmt.Println(courses)

}