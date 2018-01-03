package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	for {
		showMenu()
		input := readInput()

		switch input {
		case 1:
			triggerTest()
		case 2:
			showLogs()
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("No such command")
			os.Exit(-1)
		}
	}
}

func showMenu() {
	fmt.Println("")
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
			fmt.Println(retrieveTime(), "Availability testing:", i, website)
			testWebsite(website)
		}
		time.Sleep(5 * time.Second)
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
		fmt.Println(retrieveTime(), url, "OK")
		writeFile(url, "OK")
	} else {
		fmt.Println(retrieveTime(), url, "FAILED")
		writeFile(url, "FAILED")
	}
}

func writeFile(url string, status string) {
	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	file.WriteString(retrieveTime() + " - " + url + " - " + status + "\n")
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

func showLogs() {
	file, err := ioutil.ReadFile("logs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println(string(file))
}

func retrieveTime() string {
	return time.Now().Format("02/01/2006 15:04:05")
}
