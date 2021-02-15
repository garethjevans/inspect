package inspect_test

import (
	"fmt"
	"testing"

	"github.com/garethjevans/inspect/pkg/inspect"
	"github.com/stretchr/testify/assert"
)

func TestBaseURL(t *testing.T) {
	tests := []struct {
		labels   map[string]string
		expected string
	}{
		{
			labels: map[string]string{
				"invalid.label": "value",
			},
			expected: "",
		},
		{
			labels: map[string]string{
				"org.opencontainers.image.url": "https://github.com/org/repo.git",
			},
			expected: "https://github.com/org/repo",
		},
		{
			labels: map[string]string{
				"org.label-schema.url": "https://github.com/org/repo2.git",
			},
			expected: "https://github.com/org/repo2",
		},
		{
			labels: map[string]string{
				"org.opencontainers.image.url": "https://github.com/org/repo.git",
				"org.label-schema.url":         "https://github.com/org/repo2.git",
			},
			expected: "https://github.com/org/repo",
		},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s", tc.labels), func(t *testing.T) {
			actual := inspect.BaseURL(tc.labels)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestGitHubURL(t *testing.T) {
	tests := []struct {
		labels   map[string]string
		expected string
	}{
		{
			labels: map[string]string{
				"invalid.label": "value",
			},
			expected: "",
		},
		{
			labels: map[string]string{
				"org.opencontainers.image.url":      "https://github.com/org/repo.git",
				"org.opencontainers.image.revision": "rev1",
			},
			expected: "https://github.com/org/repo/tree/rev1",
		},
		{
			labels: map[string]string{
				"org.opencontainers.image.url":      "https://bitbucket.com/org/repo.git",
				"org.opencontainers.image.revision": "rev1",
			},
			expected: "",
		},
		{
			labels: map[string]string{
				"org.label-schema.url":     "https://github.com/org/repo2.git",
				"org.label-schema.vcs-ref": "rev2",
			},
			expected: "https://github.com/org/repo2/tree/rev2",
		},
		{
			labels: map[string]string{
				"org.opencontainers.image.url":      "https://github.com/org/repo.git",
				"org.label-schema.url":              "https://github.com/org/repo2.git",
				"org.opencontainers.image.revision": "rev1",
				"org.label-schema.vcs-ref":          "rev2",
			},
			expected: "https://github.com/org/repo/tree/rev1",
		},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s", tc.labels), func(t *testing.T) {
			actual := inspect.GitHubURL(tc.labels)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestSourceURL(t *testing.T) {
	tests := []struct {
		labels   map[string]string
		expected string
	}{
		{
			labels: map[string]string{
				"invalid.label": "value",
			},
			expected: "",
		},
		{
			labels: map[string]string{
				"org.opencontainers.image.url": "https://github.com/org/repo.git",
			},
			expected: "https://github.com/org/repo.git",
		},
		{
			labels: map[string]string{
				"org.label-schema.url": "https://github.com/org/repo2.git",
			},
			expected: "https://github.com/org/repo2.git",
		},
		{
			labels: map[string]string{
				"org.opencontainers.image.url": "https://github.com/org/repo.git",
				"org.label-schema.url":         "https://github.com/org/repo2.git",
			},
			expected: "https://github.com/org/repo.git",
		},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s", tc.labels), func(t *testing.T) {
			actual := inspect.SourceURL(tc.labels)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestRevision(t *testing.T) {
	tests := []struct {
		labels   map[string]string
		expected string
	}{
		{
			labels: map[string]string{
				"invalid.label": "value",
			},
			expected: "",
		},
		{
			labels: map[string]string{
				"org.opencontainers.image.revision": "rev1",
			},
			expected: "rev1",
		},
		{
			labels: map[string]string{
				"org.label-schema.vcs-ref": "rev2",
			},
			expected: "rev2",
		},
		{
			labels: map[string]string{
				"org.opencontainers.image.revision": "rev1",
				"org.label-schema.vcs-ref":          "rev2",
			},
			expected: "rev1",
		},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s", tc.labels), func(t *testing.T) {
			actual := inspect.Revision(tc.labels)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
