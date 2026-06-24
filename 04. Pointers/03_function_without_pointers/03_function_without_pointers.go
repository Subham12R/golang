package main

import "fmt"

func calculatePriceAfterDiscount(coffeePrice float64, discountRate float64) float64 {
	return coffeePrice - (coffeePrice * discountRate)
} 

func main() {
	var coffeePrice float64 = 5
	fmt.Printf("Basic coffee price: $%.2f\n", coffeePrice)
	coffeePrice = calculatePriceAfterDiscount(coffeePrice, 0.10)
	fmt.Printf("Final Price: $%.2f\n", coffeePrice)
}