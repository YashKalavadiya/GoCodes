package main

import "fmt"
func main() {
	fmt.Println("Structs in golang")
	//no inheritance in golang; no super or parent


	yash := User{"Yash","yashkalavadiya1010@gmail.com",true,19}

	fmt.Println(yash)

	fmt.Printf("Yash details are: %+v\n", yash)

	yash.GetStatus()

	yash.NewMail()
	fmt.Println("New Email of this user is : ", yash.Email)

}

type User struct {
	Name string
	Email string
	Status bool
	Age int
	
}

func (u User) GetStatus(){
	fmt.Println("User Status: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "new@mail.dev";
	fmt.Println("New Email of this user is : ", u.Email)
}