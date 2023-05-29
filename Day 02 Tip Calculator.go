package main

import (
	"fmt"
)

func main() {
	var Total_Bill float64
	var Tip_Percent int
	var Number_of_People int

	fmt.Println("Welcome to the tip calculator. ")
	fmt.Println("What was the total bill?: $ ")
	fmt.Scan(&Total_Bill)
	fmt.Println("What percentage tip would you like to give? 10, 12, or 15? ")
	fmt.Scan(&Tip_Percent)
	fmt.Println("How many people to split the bill? ")
	fmt.Scan(&Number_of_People)

	Tip := Total_Bill * (float64(Tip_Percent) / 100)
	Total := Total_Bill + Tip
	Total_per_person := Total / float64(Number_of_People)

	fmt.Println("Each person should pay: $", Total_per_person)
}
