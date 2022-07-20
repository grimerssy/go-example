package auth

import (
	"regexp"
)

func shouldSkip(method string) bool {
	const skipRegex = `^/v\d+.AuthService/.+`
	return regexp.
		MustCompile(skipRegex).
		MatchString(method)
}
