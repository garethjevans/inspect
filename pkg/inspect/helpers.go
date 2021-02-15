package inspect

import (
	"fmt"
	"strings"
)

// SourceURL Gets the SourceURL for the revision.
func SourceURL(labels map[string]string) string {
	return first(labels, "org.opencontainers.image.url", "org.label-schema.url")
}

// GitHubURL Gets the GitHubUrl for the revision.
func GitHubURL(labels map[string]string) string {
	return fmt.Sprintf("%s/tree/%s", BaseURL(labels), Revision(labels))
}

// BaseURL Gets the base source url without the .git suffix.
func BaseURL(labels map[string]string) string {
	gitURL := SourceURL(labels)
	if strings.HasSuffix(gitURL, ".git") {
		gitURL = strings.TrimSuffix(gitURL, ".git")
	}
	return gitURL
}

// Revision get the commit revision for this image.
func Revision(labels map[string]string) string {
	return first(labels, "org.opencontainers.image.revision", "org.label-schema.vcs-ref")
}

func first(labels map[string]string, names ...string) string {
	for _, n := range names {
		r := labels[n]
		if r != "" {
			return r
		}
	}
	return ""
}
