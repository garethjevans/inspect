package cmd

import (
	"strings"
)

// ParseRepo splits a full repository reference into repo and tag.
func ParseRepo(in string) (string, string, error) {
	if !strings.Contains(in, ":") {
		return validate(in, "latest")
	}

	// look for sha first
	if strings.Contains(in, "@") {
		parts := strings.SplitN(in, "@", 2)
		return validate(parts[0], parts[1])
	}

	parts := strings.SplitN(in, ":", 2)
	if parts[1] == "" {
		return validate(parts[0], "latest")
	}

	return validate(parts[0], parts[1])
}

func validate(repository string, tag string) (string, string, error) {
	//if strings.Count(repository, "/") > 1 {
	//	if !strings.HasPrefix(repository, "docker.io/") {
	//		return "", tag, fmt.Errorf("image hosted at registry %s not supported", repository)
	//	}
	//}
	//if !strings.Contains(repository, "/") {
	//	return "library/" + repository, tag, nil
	//}
	return repository, tag, nil
}
