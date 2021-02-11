package cmd

import "strings"

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

func Contains(in []string, val string) (ret bool) {
	for _, i := range in {
		if i == val {
			ret = true
		}
	}
	return
}
