package main

import "fmt"

func main() {
	name := "Americano"
	ready := true
	price := 3.99
	orderedCount := 5
	var stockCount int64 = 5000

	fmt.Printf("Type of name is: %T", name)

	fmt.Printf("Type of  name is :%T\n", name)
	fmt.Printf("Type of  ready is:%T\n", ready)
	fmt.Printf("Type of  price is: %T\n", price)
	fmt.Printf("Type of  count is :%T\n", orderedCount)
	fmt.Printf("Type of  count is :%T\n", stockCount)

}
