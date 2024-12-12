package main

import "fmt"

func main() {
	var a int
	var b int
	var sum int

	fmt.Println("Enter first numbers: ")
	fmt.Scan(&a)

	fmt.Println("Enter second number: ")
	fmt.Scan(&b)

	sum = a + b

	fmt.Println("Total is: ", sum)
}
