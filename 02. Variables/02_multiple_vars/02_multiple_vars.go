package main

import "fmt"

// var price float32= 2.50

func main() {
	// Explicitly declared
	var coffeeName string = "Espresso"
	// Type inferred
	var size = "Medium"
	
	
	//Short declaration and init (this type only possible in func)
	price := 2.60
	
	fmt.Println(size, coffeeName, "price is $", price)
	//String format using printf
	fmt.Printf("%s %s price is $%.2f. \n", coffeeName, size, price)
}