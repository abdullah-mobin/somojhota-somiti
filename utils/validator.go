package utils

import "regexp"

var bdPhoneRegex = regexp.MustCompile(`^(?:\+8801|8801|01)[3-9]\d{8}$`)

func IsValidBDPhoneNumber(phone string) bool {
	return bdPhoneRegex.MatchString(phone)
}
