package utils

import "unicode"

func VerifyPassword(s string) bool {
	if len(s) < 8 || len(s) > 40 {
		return false
	}

	number, upper, special := false, false, false

	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
		case isSpecialCharacter(c):
			special = true
		case number && upper && special:
			break
		}
	}

	return number && upper && special
}

func isSpecialCharacter(char rune) bool {
	return !unicode.IsLetter(char) && !unicode.IsNumber(char) && !unicode.IsSpace(char)
}
