package paths

import (
	"strings"
)

func Parse(s string) []string {
	if s == "" {
		return []string{}
	}
	if s == "/" {
		return []string{}
	}

	// Remove leading slash
	if s[0] == '/' {
		s = s[1:]
	}

	// Remove trailing slash
	if s[len(s)-1] == '/' {
		s = s[:len(s)-1]
	}
	if s == "" {
		return []string{}
	}

	parts := strings.Split(s, "/")
	return parts
}
