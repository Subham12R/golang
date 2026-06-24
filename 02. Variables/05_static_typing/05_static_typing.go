package main

import "fmt"

func main(){
	//Explicit type declaration
	var cupsQty int = 3

	fmt.Println("Number of cups: ", cupsQty)
	//Implicit Declaration
	var wasProcessed = true

	//wasProcessed = "no" // X Compile Time error
	fmt.Println("Order was processed:", wasProcessed)
}