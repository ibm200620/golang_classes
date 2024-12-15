package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	// Check if age is greater than 18
	if age >= 18 {
		fmt.Println(name, "is eligible to get a driving license.")
	} else {
		fmt.Println(name, "is not eligible to get a driving license.")
	}
}
