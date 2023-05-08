// Created by: Dominic M.
// Created on: April 2023
//
// This program multiplies two numbers together.
package main

import (
	"fmt"
)

func main() {

	var sideA int
	var sideB int
	var counter int = 0
	var output int = 0

	// input
	fmt.Println("This program multiplies two numbers together.")
	fmt.Println()
	fmt.Print("Enter your first number: ")
	fmt.Println()
	fmt.Scanln(&sideA)
	fmt.Println()
	fmt.Print("Enter your second number: ")
	fmt.Println()
	fmt.Scanln(&sideB)

	// process
	for counter < sideA {
		output = output + sideB
		counter = counter + 1
	}

	fmt.Println()
	fmt.Print(sideA, " x ", sideB, " = ", output)
	fmt.Println()
	fmt.Print("\nDone.")
}
