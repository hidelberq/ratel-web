package util

import (
	"time"
)

const HtmlLocalTimeFormat = "2006-01-02T15:04"

func ParseLocateTimeInJST(value string) (time.Time, error) {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return time.Time{}, err
	}

	return time.ParseInLocation(HtmlLocalTimeFormat, value, loc)
}
