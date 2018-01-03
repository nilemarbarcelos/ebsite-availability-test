package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	showMenu()
	input := readInput()

	switch input {
	case 1:
		triggerTest()
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

func triggerTest() {
	websites := readFile()
	for i := 0; i < 5; i++ {
		for i, website := range websites {
			fmt.Println("Availability testing:", i, website)
			testWebsite(website)
		}
		fmt.Println("")
	}
}

func testWebsite(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	if resp.StatusCode == 200 {
		fmt.Println(url, "OK")
	} else {
		fmt.Println(url, "FAILED")
	}
}

func readFile() []string {
	var websites []string
	file, err := os.Open("websites.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		websites = append(websites, line)
		if err == io.EOF {
			break
		}
	}
	file.Close()
	return websites
}
