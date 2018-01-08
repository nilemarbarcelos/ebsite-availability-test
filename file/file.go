package file

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/nilemarbarcelos/website-availability-test/timeutil"
)

func ReadFile() []string {
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

func WriteFile(url string, status string) {
	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	file.WriteString(timeutil.RetrieveTime() + " - " + url + " - " + status + "\n")
}

func ShowLogs() {
	file, err := ioutil.ReadFile("logs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	content := string(file)

	if len(content) == 0 {
		fmt.Println("Log file is empty")
	} else {
		fmt.Println(content)
	}
}
