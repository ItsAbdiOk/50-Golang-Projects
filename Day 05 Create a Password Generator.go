package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var nr_letters int
	var nr_symbols int
	var nr_numbers int
	letters := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	numbers := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	symbols := []rune{'!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '=', '+', '[', ']', '{', '}', '|', ';', ':', '"', '<', '>', ',', '.', '/', '?', '`', '~'}
	fmt.Println("Welcome to Password Generator")
	fmt.Println("How many letters would you like in your password?")
	fmt.Scan(&nr_letters)
	fmt.Println("How many symbols would you like?")
	fmt.Scan(&nr_symbols)
	fmt.Println("How many numbers would you like?")
	fmt.Scan(&nr_numbers)
	Size_of_password := nr_letters + nr_symbols + nr_numbers

	Password := make([]rune, Size_of_password)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < nr_letters; i++ {
		Password[i] = letters[rand.Intn(len(letters))]
	}
	for i := nr_letters; i < nr_letters+nr_symbols; i++ {
		Password[i] = symbols[rand.Intn(len(symbols))]
	}
	for i := nr_letters + nr_symbols; i < Size_of_password; i++ {
		Password[i] = numbers[rand.Intn(len(numbers))]
	}
	fmt.Println("Your password is:", string(Password))
}
