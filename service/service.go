package service

import (
	"regexp"
)

func isValidNumber(s string) bool {
	numericRegex := regexp.MustCompile("^[0-9]+$")
	return numericRegex.MatchString(s)
}

func countdigits(number int) int {
	count := 0
	for number > 0 {
		count++
		number = number / 10
	}
	return int(count)
}

func Validatetoken(token, SecretKey string) bool {
	_, err := ExtractCustomerID(token, SecretKey)
	return err == nil
}
