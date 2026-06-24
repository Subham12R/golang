package main

import "fmt"

func main() {
	var coffeePrice = 4.50

	fmt.Println("Coffee price: ", coffeePrice)
	fmt.Println("Memory address:", &coffeePrice)

	var pointerToCoffeePrice *float64 = &coffeePrice

	*pointerToCoffeePrice = 5.50

	fmt.Println("Updated Price: ", coffeePrice)
	fmt.Println("Memory address: ", pointerToCoffeePrice)

}