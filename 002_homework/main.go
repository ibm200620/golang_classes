package main

import "fmt"

func main() {
	var name string
	var age float32

	fmt.Println("Enter your name")
	fmt.Scan(&name)

	fmt.Println("Enter your age")
	fmt.Scan(&age)

	fmt.Println("Hello", name, "Today your are", age)
}
