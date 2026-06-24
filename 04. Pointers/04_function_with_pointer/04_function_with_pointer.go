package main

import "fmt"

func applyDiscount(coffeePrice *float64, discountRate float64)  {
	 *coffeePrice = *coffeePrice - (*coffeePrice * discountRate)
} 

func main() {
	var coffeePrice float64 = 5.00
	fmt.Println("Memory location of price: ", &coffeePrice)
	var discountRate float64 = 0.10

	fmt.Printf("Basic coffee price: $%.2f\n", coffeePrice)
	applyDiscount(&coffeePrice, discountRate)
	fmt.Printf("Final Price: $%.2f\n", coffeePrice)
}