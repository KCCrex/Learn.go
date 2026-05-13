package main

import "fmt"

func main() {

	var (
		coffeeType string  = "Latte"
		quantity   int     = 3
		price      float64 = 4.25
	)

	fmt.Printf("coffeeType: %s, quantity: %d, price: %.2f\n", coffeeType, quantity, price)

	var (
		customerName string = "Isa"
		tableNum     int    = 5
		isReadyToPay bool   = false
	)
	fmt.Printf("customerName %s at table %d , is ready to pay: %t", customerName, tableNum, isReadyToPay)
	const (
		sizeSmall  = "S"
		sizeMedium = "M"
		sizeLarge  = "L"
	)
}
