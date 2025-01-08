package main

import "fmt"

func main() {
	var username string = "Ayush"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	if isLoggedIn {
		fmt.Printf("User is Logged In")
	}else {
		fmt.Printf("User is Not Logged In")
	}

	var smallValue int = 9801585
	fmt.Println(smallValue)
}