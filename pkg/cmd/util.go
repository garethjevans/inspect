package cmd

import "strings"

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

// AllKeys returns a list of all unique keys from all maps.
func AllKeys(maps ...map[string]string) []string {
	keys := []string{}
	for _, m := range maps {
		for k := range m {
			if !Contains(keys, k) {
				keys = append(keys, k)
			}
		}
	}
	return keys
}

// Contains returns true if the string is contained within this list.
func Contains(in []string, val string) (ret bool) {
	for _, i := range in {
		if i == val {
			ret = true
		}
	}
	return
}
