package util

import (
	"log"
	"time"
)

const FormatDateTimeLocalPCChrome = "2006-01-02T15:04"
const FormatDateTimeLocaliPhoneSafari = "2006-01-02T15:04:05.999"

func ParseLocateTimeInJST(value string) (time.Time, error) {
	log.Println(value)

	layouts := []string{
		time.RFC3339,
		time.RFC3339Nano,
		FormatDateTimeLocalPCChrome,
		FormatDateTimeLocaliPhoneSafari,
	}

	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return time.Time{}, err
	}

	for _, l := range layouts {
		t, err := time.ParseInLocation(l, value, loc)
		if err == nil {
			return t, nil
		}
	}

	return time.Time{}, nil
}
