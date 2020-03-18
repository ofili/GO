package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Guess the Number!")
	fmt.Println("***************************")

	// geneate a random number
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(100) // generates number between 0 and n (10)

	var guess int

	for {
		fmt.Println("Please enter a number between 1 and 100 : ")
		fmt.Scan(&guess)

		if guess > secretNumber {
			fmt.Println("Too Big, try again")
		} else if guess < secretNumber {
			fmt.Println("Too Small, try again")
		} else {
			fmt.Println("Correct. You Win!")
			break
		}
	}

}
