package main

//program to count the lines in a file

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//get file from command line
	filePath := os.Args[1]
	//open and read the file
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)

	}
	//scan the file
	fileScanner := bufio.NewScanner(readFile)
	//set lineCount to start at 0
	lineCount := 0
	//for loop to run through the file
	for fileScanner.Scan() {
		lineCount++
	}

	fmt.Println("number of lines:", lineCount)

}
