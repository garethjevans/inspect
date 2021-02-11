package cmd_test

import (
	"testing"

	"github.com/garethjevans/inspect/pkg/cmd"
	"github.com/stretchr/testify/assert"
)

func TestCanParseArgs(t *testing.T) {
	tests := []struct {
		input         string
		expectedRepo  string
		expectedTag   string
		expectedError bool
	}{
		{
			input:        "jenkins/jenkins:jdk11-hotspot-windowsservercore-2019",
			expectedRepo: "jenkins/jenkins",
			expectedTag:  "jdk11-hotspot-windowsservercore-2019",
		},
		{
			input:        "jenkins/jenkins:jdk11-hotspot-windowsservercore-1809",
			expectedRepo: "jenkins/jenkins",
			expectedTag:  "jdk11-hotspot-windowsservercore-1809",
		},
		{
			input:        "jenkins/jenkins@sha256:1234567890",
			expectedRepo: "jenkins/jenkins",
			expectedTag:  "sha256:1234567890",
		},
		{
			input:        "jenkins/jenkins",
			expectedRepo: "jenkins/jenkins",
			expectedTag:  "latest",
		},
		{
			input:        "alpine:3.13.0",
			expectedRepo: "library/alpine",
			expectedTag:  "3.13.0",
		},
		{
			input:        "alpine",
			expectedRepo: "library/alpine",
			expectedTag:  "latest",
		},
		{
			input:         "gcr.io/random/image:1",
			expectedError: true,
		},
		{
			input:         "quay.io/random/image:1",
			expectedError: true,
		},
		{
			input:         "other.io/random/image:1",
			expectedError: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			repo, tag, err := cmd.ParseRepo(tc.input)
			if tc.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedRepo, repo)
				assert.Equal(t, tc.expectedTag, tag)
			}
		})
	}
}
