package cmd

import (
	"strings"
)

// ParseRepo splits a full repository reference into repo and tag.
func ParseRepo(in string) (string, string) {
	if !strings.Contains(in, ":") {
		return in, "latest"
	}

	// look for sha first
	if strings.Contains(in, "@") {
		parts := strings.SplitN(in, "@", 2)
		return parts[0], parts[1]
	}

	parts := strings.SplitN(in, ":", 2)
	if parts[1] == "" {
		return parts[0], "latest"
	}

	return parts[0], parts[1]
}
