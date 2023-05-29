package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	rock = `
    _______
---'   ____)
      (_____)
      (_____)
      (____)
---.__(___)
`
	paper = `
    _______
---'   ____)____
          ______)
          _______)
         _______)
---.__________)
`
	scissors = `
    _______
---'   ____)____
          ______)
       __________)
      (____)
---.__(___)
`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Print("What do you choose? Type 0 for Rock, 1 for Paper, or 2 for Scissors: ")
	var userChoice int
	_, err := fmt.Scan(&userChoice)
	if err != nil {
		fmt.Println("Invalid input. Please try again.")
		return
	}

	switch userChoice {
	case 0:
		fmt.Println(rock)
		fmt.Println("You chose rock")
	case 1:
		fmt.Println(paper)
		fmt.Println("You chose paper")
	case 2:
		fmt.Println(scissors)
		fmt.Println("You chose scissors")
	default:
		fmt.Println("Invalid choice. Please try again.")
		return
	}

	computerChoice := rand.Intn(3)
	switch computerChoice {
	case 0:
		fmt.Println(rock)
		fmt.Println("Computer chose rock")
	case 1:
		fmt.Println(paper)
		fmt.Println("Computer chose paper")
	case 2:
		fmt.Println(scissors)
		fmt.Println("Computer chose scissors")
	}

	switch {
	case userChoice < 0 || userChoice >= 3:
		fmt.Println("You typed an invalid number. You lose!")
	case userChoice == 0 && computerChoice == 2:
		fmt.Println("You win!")
	case computerChoice == 0 && userChoice == 2:
		fmt.Println("You lose")
	case computerChoice > userChoice:
		fmt.Println("You lose")
	case userChoice > computerChoice:
		fmt.Println("You win!")
	case computerChoice == userChoice:
		fmt.Println("It's a draw")
	}
}
