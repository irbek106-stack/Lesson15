package main

import "fmt"

func main(){
	user1 := Reader{
		ID: 1,
		FirstName: "Tofl",
		LastName: "Gamigoev",
		IsActive: true,
	}

	fmt.Println(user1)
	user1.DisplayReader()
	user1.Deactivate()
	user1.Activate()
	fmt.Println(user1)
}