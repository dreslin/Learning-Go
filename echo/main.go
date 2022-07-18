package main

import (
	"fmt"
	"os"
)

func main() {
	echo := os.Args[0]
	cmdArgs := os.Args[1:]
	index := 0
	for _, cmdArg := range cmdArgs {
		fmt.Println(index, cmdArg[index])
	}
	fmt.Printf("%#v", cmdArgs)
	fmt.Println(cmdArgs)
	fmt.Println(len(cmdArgs))
	fmt.Println(echo)

}
