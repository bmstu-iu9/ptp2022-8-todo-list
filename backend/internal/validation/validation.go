package validation

import "regexp"

// ValidateField performs basic validation on string field.
func ValidateField(field string, minLen, maxLen int, regex string) bool {
	if matched, _ := regexp.MatchString(regex, field); !matched {
		return false
	}
	if len(field) < minLen || len(field) > maxLen {
		return false
	}
	return true
}
