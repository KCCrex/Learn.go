package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1
	var guess int
	attempts := 0
	fmt.Printf("Welcome to the Number Guessing Game!")
	fmt.Println(" 'm thinking of a number between 1 and 100. Can you guess it?")
	for {
		attempts++
		fmt.Print("Enter your guess:")
		_, err := fmt.Scan(&guess)

		if err != nil {
			fmt.Println("Please enter a valid number!")
			continue
		}
		if guess < target {
			fmt.Println("Oops! Too Low! Try Again !")
		} else if guess > target {
			fmt.Println("Too High! Try Again!")
		} else {
			fmt.Printf("You got it ! The number was %d ,I took you %d attempts to guess the number,Congratulations \n!", target, attempts)

		}

	}
}
