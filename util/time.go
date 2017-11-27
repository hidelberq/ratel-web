package util

import (
	"log"
	"time"
)

const FormatDateTimeLocal = "2006-01-02T15:04"

func ParseLocateTimeInJST(value string) (time.Time, error) {
	log.Println(value)
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return time.Time{}, err
	}

	return time.ParseInLocation(FormatDateTimeLocal, value, loc)
}
