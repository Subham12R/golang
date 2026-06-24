package main

import "fmt"

func main() {
	coffee := "Espresso"
	pointer := &coffee

	fmt.Println("Coffee name : ", coffee)
	fmt.Println("Memory location:", pointer)
	fmt.Printf("Pointer address: %p\n", &pointer)

	//Initializing copy pointer

	coffeePointer := coffee;

	fmt.Println("Coffee name : ", coffeePointer)
	fmt.Println("Memory location:", &coffeePointer)
	fmt.Printf("Pointer address: %p\n", &coffeePointer)

}