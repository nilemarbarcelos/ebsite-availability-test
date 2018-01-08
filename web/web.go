package web

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/nilemarbarcelos/website-availability-test/file"
	"github.com/nilemarbarcelos/website-availability-test/timeutil"
)

func TriggerTest() {
	websites := file.ReadFile()
	for i := 0; i < 5; i++ {
		for i, website := range websites {
			fmt.Println(timeutil.RetrieveTime(), "Availability testing:", i, website)
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
		fmt.Println(timeutil.RetrieveTime(), url, "OK")
		file.WriteFile(url, "OK")
	} else {
		fmt.Println(timeutil.RetrieveTime(), url, "FAILED")
		file.WriteFile(url, "FAILED")
	}
}
