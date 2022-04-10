package main

import "fmt"

const LoginToken string = "bxasdajhebs"

func main() {
	var username string = "Yash Manish Bina Kalavadiya"
	
	fmt.Println(username)

	fmt.Printf("Variable is of type: %T \n ", username)
	
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n ", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 321.213555555555555555555
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)


	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)


	//Implicit type
	var website = "learncodeonline.in"
	fmt.Println(website)


	//no var style

	numberOfUser := 3000000
	fmt.Println(numberOfUser)

	//printing global const
	fmt.Println(LoginToken)
	fmt.Printf("Variable type is : %T \n", LoginToken)

}