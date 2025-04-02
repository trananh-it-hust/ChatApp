package util

import "regexp"

func IsValidEmail(email string) bool {
	if len(email) < 5 || len(email) > 50 {
		return false
	}
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(regex, email)
	return matched
}
