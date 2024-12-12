package main

import "fmt"

func main() {
	var name string
	fmt.Println("Enter your name: ")
	fmt.Scan(&name)

	var age int
	fmt.Println("Enter your age: ")
	fmt.Scan(&age)

	var height float32
	fmt.Println("Enter your height: ")
	fmt.Scan(&height)

	fmt.Println("Name: ", name, age)
	fmt.Println("Age: ", age)
	fmt.Println("Height in feet: ", height)
}

// name = "Dave"
// age = 25
// height = 5.11
