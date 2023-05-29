package main

import (
	"fmt"
)

func main() {
	var City string
	var Pet string
	fmt.Println("Welcome to the Band name generator. ")
	fmt.Println("What's the name of the city you grew up in? ")
	fmt.Scan(&City)
	fmt.Println("What's your pet's name? ")
	fmt.Scan(&Pet)
	fmt.Printf("Your band name could be %s %s", City, Pet)
}
