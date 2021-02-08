package pkg

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/garethjevans/inspect/pkg/inspect"

	"github.com/stretchr/testify/assert"
)

func TestCanQueryDockerManifest(t *testing.T) {
	// get token
	repo := "jenkinsciinfra/terraform"
	tag := "latest"

	client := inspect.Client{
		Client: http.Client{},
	}

	l, err := client.Labels(repo, tag)
	assert.NoError(t, err)

	t.Log(l)

	// https://github.com/jenkins-infra/docker-terraform/tree/d25f040

	gitURL := l["org.opencontainers.image.source"]
	if strings.HasSuffix(gitURL, ".git") {
		gitURL = strings.TrimSuffix(gitURL, ".git")
	}

	t.Log(fmt.Sprintf("%s/tree/%s", gitURL, l["org.opencontainers.image.revision"]))
}
