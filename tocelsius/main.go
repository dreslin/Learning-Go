// converts temperature from Fahrenheit to Celsius.
package main

import (
	"fmt"

	"github.com/dreslin/Learning-Go/keyboard"
)

func main() {
	fmt.Print("Enter a temperature in Fahrenheit: ")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		panic(err)
	}

	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%0.2f degrees Celsius\n", celsius)

}
