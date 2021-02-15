package cmd_test

import (
	"testing"

	"github.com/garethjevans/inspect/pkg/cmd"
	"github.com/stretchr/testify/assert"
)

func TestCanParseArgs(t *testing.T) {
	tests := []struct {
		input        string
		expectedRepo string
		expectedTag  string
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
			expectedRepo: "alpine",
			expectedTag:  "3.13.0",
		},
		{
			input:        "alpine",
			expectedRepo: "alpine",
			expectedTag:  "latest",
		},
		{
			input:        "gcr.io/random/image:1",
			expectedRepo: "gcr.io/random/image",
			expectedTag:  "1",
		},
		{
			input:        "quay.io/random/image:1",
			expectedRepo: "quay.io/random/image",
			expectedTag:  "1",
		},
		{
			input:        "other.io/random/image:1",
			expectedRepo: "other.io/random/image",
			expectedTag:  "1",
		},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			repo, tag := cmd.ParseRepo(tc.input)
			assert.Equal(t, tc.expectedRepo, repo)
			assert.Equal(t, tc.expectedTag, tag)
		})
	}
}
