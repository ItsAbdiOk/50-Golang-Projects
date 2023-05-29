package main

import (
	"fmt"
)

func main() {
	// Declare variables
	var Q1 string
	var Q2 string
	var Q3 string

	fmt.Println("Welcome to Treasure Island. Your mission is to find the treasure.")
	fmt.Println("You're at a cross road. Where do you want to go? Type \"left\" or \"right\".")
	fmt.Scan(&Q1)
	if Q1 == "left" {
		fmt.Println("You come to a lake. There is an island in the middle of the lake. Type \"wait\" to wait for a boat. Type \"swim\" to swim across.")
		fmt.Scan(&Q2)
		if Q2 == "wait" {
			fmt.Println("You arrive at the island unharmed. There is a house with 3 doors. One red, one yellow, and one blue. Which color do you choose?")
			fmt.Scan(&Q3)
			if Q3 == "red" {
				fmt.Println("It's a room full of fire. Game Over.")
			} else if Q3 == "yellow" {
				fmt.Println("You found the treasure! You Win!")
			} else if Q3 == "blue" {
				fmt.Println("You enter a room of beasts. Game Over.")
			} else {
				fmt.Println("You chose a door that doesn't exist. Game Over.")
			}
		} else if Q2 == "swim" {
			fmt.Println("You get attacked by an angry trout. Game Over.")
		} else {
			fmt.Println("You chose a path that doesn't exist. Game Over.")
		}
	} else if Q1 == "right" {
		fmt.Println("You fell into a hole. Game Over.")
	} else {
		fmt.Println("You chose a path that doesn't exist. Game Over.")
	}
}
