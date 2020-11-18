package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
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

func CheckJwtExpire(exp interface{}) bool {
	strExp := fmt.Sprintf("%v", exp)
	expire, _ := time.Parse(time.RFC3339, strExp)
	diff := expire.Sub(time.Now())
	if diff <= 0 {
		return false
	}
	return true
}
