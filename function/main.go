//always start with this
package main

//import required packages
import (
	"fmt"
)

//declare and define function and variables
func double(number int) {
	//'return' must be last line in function, program exits at 'return' line
	number *= 2
}

//Call function and play
func main() {
	dozen := double(6.0)
	fmt.Println(dozen)
	fmt.Println(double(4.2))
}
