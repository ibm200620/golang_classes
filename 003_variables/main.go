package main

import (
	"fmt"
)

func main() {
	var message string // Declaring a variable

	message = "It's Saturday" // Initializing a variable

	fmt.Println(message)

	var a int

	a = 5

	fmt.Println(a)

	var b int
	fmt.Println(b)

	var itemName string
	fmt.Println(itemName)

	var price float32
	price = 39.9512546544665544
	fmt.Println(price)

	var inStock bool
	inStock = false
	fmt.Println(inStock)

	var alphabet byte
	alphabet = 'A'
	fmt.Println(alphabet)

	// rune for a chinese character
	var alpha rune
	alpha = '中'
	fmt.Println(alpha)
}

// Every variable in Go has a type
// A variable needs to be declared
// A variable can be initialized
