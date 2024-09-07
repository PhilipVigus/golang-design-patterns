package main

import "fmt"

func main() {
	println("Hello, World!")

	var explicit string = "Explicit"

	implicit := "Implicit"

	println(explicit)
	println(implicit)

	my_array := [5]int{1, 2, 3, 4, 5}
	for index, value := range my_array {
		fmt.Printf("Index is %d and value is %d", index, value)
	}
}
