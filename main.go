package main

import (
	"fmt"
	"os"

	"github.com/nilemarbarcelos/website-availability-test/file"
	"github.com/nilemarbarcelos/website-availability-test/input"
	"github.com/nilemarbarcelos/website-availability-test/menu"
	"github.com/nilemarbarcelos/website-availability-test/web"
)

func main() {
	for {
		menu.ShowMenu()
		input := input.ReadInput()

		switch input {
		case 1:
			web.TriggerTest()
		case 2:
			file.ShowLogs()
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("No such command")
			os.Exit(-1)
		}
	}
}
