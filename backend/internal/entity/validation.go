package entity

import "regexp"

func validateField(field string, minLen, maxLen int, regex string) bool {
	if matched, _ := regexp.MatchString(regex, field); !matched ||
		len(field) < minLen || len(field) > maxLen {
		return false
	}
	return true
}
