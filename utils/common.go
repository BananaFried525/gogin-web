package utils

import (
	"regexp"
	"strconv"
)

func IsEmail(email string) bool {
	r, _ := regexp.Compile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return r.MatchString(email)
}

func IsNumber(number string) bool {
	if _, err := strconv.Atoi(number); err != nil {
		return false
	}
	return true
}
