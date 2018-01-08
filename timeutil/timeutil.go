package timeutil

import "time"

const brazilianFormat = "02/01/2006 15:04:05"

func RetrieveTime() string {
	return time.Now().Format(brazilianFormat)
}
