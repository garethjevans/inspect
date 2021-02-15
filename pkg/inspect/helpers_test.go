package inspect_test

import (
	"fmt"
	"testing"

	"github.com/garethjevans/inspect/pkg/inspect"
	"github.com/stretchr/testify/assert"
)

func TestBaseURL(t *testing.T) {
	tests := []struct {
		labels          map[string]string
		expectedBaseURL string
	}{
		{
			labels: map[string]string{
				"invalid.label": "value",
			},
			expectedBaseURL: "",
		},
		{
			labels: map[string]string{
				"org.opencontainers.image.url":  "https://github.com/base/url.git",
				"org.opencontainers.image.url1": "https://github.com/base/url1.git",
				"org.opencontainers.image.url2": "https://github.com/base/url2.git",
			},
			expectedBaseURL: "https://github.com/base/url",
		},
		{
			labels: map[string]string{
				"org.label-schema.url": "https://github.com/base/url2.git",
			},
			expectedBaseURL: "https://github.com/base/url2",
		},
		{
			labels: map[string]string{
				"org.opencontainers.image.url": "https://github.com/org/repo.git",
				"org.label-schema.url":         "https://github.com/org/repo2.git",
				"org.label-schema.revision":    "doesn't matter",
			},
			expectedBaseURL: "https://github.com/org/repo",
		},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s", tc.labels), func(t *testing.T) {
			actual := inspect.BaseURL(tc.labels)
			assert.Equal(t, tc.expectedBaseURL, actual)
		})
	}
}

func TestGitHubURL(t *testing.T) {
	tests := []struct {
		labels            map[string]string
		expectedGitHubURL string
	}{
		{
			labels: map[string]string{
				"invalid.label": "value",
			},
			expectedGitHubURL: "",
		},
		{
			labels: map[string]string{
				"org.opencontainers.image.url":      "https://github.com/org/repo.git",
				"org.opencontainers.image.revision": "rev1",
			},
			expectedGitHubURL: "https://github.com/org/repo/tree/rev1",
		},
		{
			labels: map[string]string{
				"org.opencontainers.image.url":      "https://bitbucket.com/org/repo.git",
				"org.opencontainers.image.revision": "rev1",
			},
			expectedGitHubURL: "",
		},
		{
			labels: map[string]string{
				"org.label-schema.url":     "https://github.com/org/repo2.git",
				"org.label-schema.vcs-ref": "rev2",
			},
			expectedGitHubURL: "https://github.com/org/repo2/tree/rev2",
		},
		{
			labels: map[string]string{
				"org.opencontainers.image.url":      "https://github.com/org/repo.git",
				"org.label-schema.url":              "https://github.com/org/repo2.git",
				"org.opencontainers.image.revision": "rev1",
				"org.label-schema.vcs-ref":          "rev2",
			},
			expectedGitHubURL: "https://github.com/org/repo/tree/rev1",
		},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s", tc.labels), func(t *testing.T) {
			gitHubURL := inspect.GitHubURL(tc.labels)
			assert.Equal(t, tc.expectedGitHubURL, gitHubURL)
		})
	}
}

func TestSourceURL(t *testing.T) {
	tests := []struct {
		labels            map[string]string
		expectedSourceURL string
	}{
		{
			labels: map[string]string{
				"invalid.label": "value",
			},
			expectedSourceURL: "",
		},
		{
			labels: map[string]string{
				"org.opencontainers.image.url": "https://github.com/source/url.git",
			},
			expectedSourceURL: "https://github.com/source/url.git",
		},
		{
			labels: map[string]string{
				"org.label-schema.url": "https://github.com/other/url.git",
			},
			expectedSourceURL: "https://github.com/other/url.git",
		},
		{
			labels: map[string]string{
				"org.opencontainers.image.url": "https://github.com/org/repo.git",
				"org.label-schema.url":         "https://github.com/org/repo2.git",
			},
			expectedSourceURL: "https://github.com/org/repo.git",
		},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s", tc.labels), func(t *testing.T) {
			sourceURL := inspect.SourceURL(tc.labels)
			assert.Equal(t, tc.expectedSourceURL, sourceURL)
		})
	}
}

func TestRevision(t *testing.T) {
	tests := []struct {
		labels           map[string]string
		expectedRevision string
	}{
		{
			labels: map[string]string{
				"invalid.label": "value",
				"random":        "v",
			},
			expectedRevision: "",
		},
		{
			labels: map[string]string{
				"org.opencontainers.image.revision": "1234567890",
			},
			expectedRevision: "1234567890",
		},
		{
			labels: map[string]string{
				"org.label-schema.vcs-ref": "rev2",
			},
			expectedRevision: "rev2",
		},
		{
			labels: map[string]string{
				"org.opencontainers.image.revision": "rev1",
				"org.label-schema.vcs-ref":          "1234567890",
			},
			expectedRevision: "rev1",
		},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s", tc.labels), func(t *testing.T) {
			revision := inspect.Revision(tc.labels)
			assert.Equal(t, tc.expectedRevision, revision)
		})
	}
}
