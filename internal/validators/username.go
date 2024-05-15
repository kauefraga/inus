package validators

import "strings"

func IsUserNameInvalid(name string) bool {
	hasUsernameInvalidlength := len(name) < 4 || len(name) > 50
	isUsernameEmpty := len(strings.TrimSpace(name)) == 0

	if hasUsernameInvalidlength || isUsernameEmpty {
		return true
	}

	return false
}
