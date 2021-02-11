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
		{"jenkins/jenkins:jdk11-hotspot-windowsservercore-2019", "jenkins/jenkins", "jdk11-hotspot-windowsservercore-2019"},
		{"jenkins/jenkins:jdk11-hotspot-windowsservercore-1809", "jenkins/jenkins", "jdk11-hotspot-windowsservercore-1809"},
		{"jenkins/jenkins@sha256:1234567890", "jenkins/jenkins", "sha256:1234567890"},
		{"jenkins/jenkins", "jenkins/jenkins", "latest"},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			repo, tag := cmd.ParseRepo(tc.input)
			assert.Equal(t, tc.expectedRepo, repo)
			assert.Equal(t, tc.expectedTag, tag)
		})
	}
}
