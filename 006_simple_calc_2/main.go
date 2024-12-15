package main

import "fmt"

func main() {
	var a float32
	var b float32
	var sum float32

	fmt.Println("Enter first numbers: ")
	fmt.Scan(&a)

	fmt.Println("Enter second number: ")
	fmt.Scan(&b)

	sum = a + b

	fmt.Println("Total is: *", sum)
}
