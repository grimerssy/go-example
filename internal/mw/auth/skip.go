package auth

import (
	"regexp"
)

const (
	skipAuthRegex = `^/v\d+.AuthService/.+`
)

func shouldSkip(method string) bool {
	return regexp.
		MustCompile(skipAuthRegex).
		MatchString(method)
}
