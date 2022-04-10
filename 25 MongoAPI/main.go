package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yash10102002/mongoapi/router"
)

func main() {
	fmt.Println("Mongodb API")
	fmt.Println("Server is getting started...")

	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))

	fmt.Println("Server is up and running on port 4000...")
	
}