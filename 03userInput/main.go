package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var welcome string = "Welcome to user input"
	fmt.Println((welcome))

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name : ")

	// comma ok || error ok
	// comma ok is used to handle two values at a time
	// error ok is used to handle one value and one error at a time
	input, _ := reader.ReadString('\n')
	fmt.Println("Welcome ", input)
	fmt.Printf("Input Type : %T",input)
}