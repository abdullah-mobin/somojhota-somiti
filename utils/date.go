package utils

import (
	"time"

	"github.com/abdullah-mobin/somojhota-somiti/config"
)

func ParseDateOnly(dateStr string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, err
	}

	location := config.AppLocation

	return time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		0, 0, 0, 0,
		location,
	), nil
}

func ParseMonthYear(dateStr string) (int, int) {
	t, _ := time.Parse("2006-01", dateStr)
	month := int(t.Month())
	year := t.Year()
	return month, year
}
