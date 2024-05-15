package validators

import "strings"

func IsEmailInvalid(email string) bool {
	isEmailEmpty := len(strings.TrimSpace(email)) == 0

	return isEmailEmpty
}
