package time

import (
	"time"
)

const (
	ISO8601 = "2006-01-02T15:04:05Z"
)

func UTCNow() string {
	return time.Now().UTC().Format(ISO8601)
}

func UTCP8Now() string {
	return time.Now().In(time.FixedZone("UTC+8", 8*60*60)).Format(ISO8601)
}

func Parse(t string) (time.Time, error) {
	var deft time.Time
	parsedTime, err := time.Parse(ISO8601, t)
	if err != nil {
		return deft, err
	}
	return parsedTime, nil
}
