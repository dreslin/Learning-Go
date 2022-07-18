// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

func main() {
	// UNCOMMENT THE CODE BELOW:

	var (
		planet string
		isTrue bool
		temp   float64
	)
	// ADD YOUR CODE BELOW:

	planet, isTrue, temp = "Mars", true, 19.5

	fmt.Println("The air is good on", planet)
	fmt.Println("It's", isTrue)
	fmt.Println("It is", temp, "degrees")

}
