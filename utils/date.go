package utils

import (
	"time"

	"github.com/abdullah-mobin/somojhota-somiti/config"
)

func ParseDateOnly(dateStr string) (string, error) {
	location := config.AppLocation
	t, err := time.ParseInLocation("2006-01-02", dateStr, location)
	if err != nil {
		return "", err
	}

	date := time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		0, 0, 0, 0,
		location,
	).Format("2006-01-02")
	return date, nil
}

func ParseMonthYear(dateStr string) (int, int, error) {
	location := config.AppLocation

	t, err := time.ParseInLocation("2006-01-02", dateStr, location)
	if err != nil {
		return 0, 0, err
	}

	return int(t.Month()), t.Year(), nil
}
