package main

import "fmt"

func main() {

	// Declare and Initialize with var with explicit type
	var coffeename string = "Espresso"
	//type inferred
	var size = "Small"
	// short declaration and initialization . Possible only inside function .
	var price = 2.50

	fmt.Println("Small Espresso price is $2.50")
	fmt.Println(size, coffeename, "price is $", price)
	fmt.Printf(" %s %s price is $%.2f\n", size, coffeename, price)

}
