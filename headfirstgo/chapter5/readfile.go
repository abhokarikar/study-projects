package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Print("Please enter name of the file: ")
	// Create a new reader to read from standard input
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	input = input[:len(input)-1] // Remove the newline character
	fmt.Println("You entered:", input)

	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a new scanner to read from a file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
