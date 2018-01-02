package main

import (
	"fmt"
	"os"
)

func main() {
	showMenu()
	input := readInput()

	switch input {
	case 1:
		fmt.Println("Testing...")
	case 2:
		fmt.Println("Showing logs...")
		os.Exit(0)
	case 0:
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("No such command")
		os.Exit(-1)
	}
}

func showMenu() {
	fmt.Println("--------------------------")
	fmt.Println("1 - Availability test")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")
	fmt.Println("--------------------------")
	fmt.Println("")
}

func readInput() int {
	var input int
	fmt.Scan(&input)
	return input
}
