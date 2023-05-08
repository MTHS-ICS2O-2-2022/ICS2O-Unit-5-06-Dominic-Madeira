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
	var counter int
	var output int

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
		output += sideB
		counter++
	}

	fmt.Println()
	fmt.Print(sideA, " x ", sideB, " = ", output)
	fmt.Println()
	fmt.Print("\nDone.")
}
