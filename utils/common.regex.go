package utils

import (
	"net"
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// the following function does not check whether the email is registered
func IsValidEmail(e string) bool {
	// too long or too short
	if len(e) < 4 || len(e) > 254 {
		return false
	}
	// forbid illegal characters
	if !emailRegex.MatchString(e) {
		return false
	}
	// is the domain valid
	parts := strings.Split(e, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return false
	}
	return true
}

func RemoveCharacter(s, chartype string) string {
	switch chartype{
	case " ":
		return strings.ReplaceAll(s, " ", "")
	case "\"":
		return strings.ReplaceAll(s, "\"", "")
	default:
		return "!!!"+s
	}
}