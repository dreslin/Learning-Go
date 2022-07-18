// RPS plays a game yada adya
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// isAnOption
func isAnOption(option string, options []string) (int, bool) {
	for i, value := range options {
		if option == value {
			return i, true
		}
	}

	return 0, false
}

func main() {
	// this can be extended, but it must be listed as a circle of an odd number of
	// items where any item defeats the next half of the circle.
	options := []string{
		"Rock",
		"Scissors",
		"Paper",
	}

	// Initializing the seed for the random number generator.
	// It uses the time now as Unix time, an integer
	// version of the current time.
	rand.Seed(time.Now().UnixNano())

	// Picks a number from 0 to 3, not including 3.
	compSeli := rand.Intn(len(options))
	compSel := options[compSeli]

	if len(os.Args) <= 1 {
		fmt.Println("---- Welcome to Rock, Papers, Scissors; The game! ----")
		fmt.Println("---- Please insert one of the following options to play: ----")
		fmt.Println(strings.Join(options, ", "))
		os.Exit(2)
	}

	userSel := strings.Title(os.Args[1])

	index, ok := isAnOption(userSel, options)
	if !ok {
		fmt.Print("---- Please enter a valid option: ----\n")
		fmt.Println(strings.Join(options, ", "))
		os.Exit(2)
	}

	if compSel == userSel {
		fmt.Printf("Computer chose %v, it's a tie!\n", compSel)
		os.Exit(0)
	}

	// Treat index as a unwrapped circle by adding the length if index is less
	// than compSeli
	if index < compSeli {
		index += len(options)
	}

	if index-compSeli <= (len(options)-1)/2 {
		fmt.Printf("Computer chose %s, you lose\n", compSel)
		os.Exit(0)
	}

	// This is the only option left
	fmt.Printf("Computer chose %s, you win\n", compSel)
}
