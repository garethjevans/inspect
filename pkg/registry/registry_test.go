package registry_test

import (
	"testing"

	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	tests := []struct {
		image string
	}{
		{
			image: "jenkins/jenkins:jdk11-hotspot-windowsservercore-2019",
		},
		{
			image: "jenkins/jenkins",
		},
		{
			image: "ghcr.io/external-secrets/kubernetes-external-secrets",
		},
		{
			image: "kiwigrid/k8s-sidecar:0.1.275",
		},
		{
			image: "jenkinsciinfra/terraform:latest",
		},
		{
			image: "jenkinsciinfra/jenkins-weekly@sha256:209c51f7b4bbc5fe2011b766344e7ea24c25b4e059573d1dd4163ac74761b785",
		},
	}
	for _, tc := range tests {
		t.Run(tc.image, func(t *testing.T) {
			ref, err := name.ParseReference(tc.image)
			assert.NoError(t, err)

			img, err := remote.Image(ref)
			assert.NoError(t, err)

			configFile, err := img.ConfigFile()
			assert.NoError(t, err)

			t.Log(configFile.Config.Labels)
		})
	}
}
