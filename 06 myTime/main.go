package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time package study")

	presentTime := time.Now()

	fmt.Println(presentTime)


	fmt.Println(presentTime.Format("01-02-2006"))

	createdDate := time.Date(2002, time.October, 10, 23, 23, 0, 0, time.UTC)

	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
}