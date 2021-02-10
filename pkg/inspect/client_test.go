package inspect_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path"
	"strings"
	"testing"

	"github.com/garethjevans/inspect/pkg/inspect"
	"github.com/garethjevans/inspect/pkg/inspect/mocks"

	"github.com/stretchr/testify/assert"
)

var (
	repo = "jenkinsciinfra/terraform"
)

func TestCanQueryDockerManifest(t *testing.T) {
	// get token
	tag := "latest"

	mock := &mocks.MockClient{}

	client := inspect.Client{
		Client: mock,
	}

	stubWithFixture(t, mock, 200, "token.json")
	stubWithFixture(t, mock, 200, "manifest.json")
	stubWithFixture(t, mock, 200, "blobs.json")

	l, err := client.Labels(repo, tag)
	assert.NoError(t, err)

	t.Log(l)

	gitURL := l["org.opencontainers.image.source"]
	if strings.HasSuffix(gitURL, ".git") {
		gitURL = strings.TrimSuffix(gitURL, ".git")
	}

	fullGitURL := fmt.Sprintf("%s/tree/%s", gitURL, l["org.opencontainers.image.revision"])
	t.Log(fullGitURL)

	assert.Equal(t, "https://github.com/jenkins-infra/docker-terraform/tree/d25f040", fullGitURL)
}

func TestCanQueryDockerManifest_UsingSha(t *testing.T) {
	// get token
	tag := "sha256:71e99d70bef50da077f0595044f3330b074e2d0417e486a12144d67cb5dd5603"

	mock := &mocks.MockClient{}

	client := inspect.Client{
		Client: mock,
	}

	stubWithFixture(t, mock, 200, "token.json")
	stubWithFixture(t, mock, 200, "manifest.json")
	stubWithFixture(t, mock, 200, "blobs.json")

	l, err := client.Labels(repo, tag)
	assert.NoError(t, err)

	t.Log(l)

	gitURL := l["org.opencontainers.image.source"]
	if strings.HasSuffix(gitURL, ".git") {
		gitURL = strings.TrimSuffix(gitURL, ".git")
	}

	fullGitURL := fmt.Sprintf("%s/tree/%s", gitURL, l["org.opencontainers.image.revision"])
	t.Log(fullGitURL)

	assert.Equal(t, "https://github.com/jenkins-infra/docker-terraform/tree/d25f040", fullGitURL)
}

func TestCanQueryDockerManifest_UsingUnknownSha(t *testing.T) {
	// get token
	tag := "sha256:71e99d70bef50da077f0595044f3330b074e2d0417e486a12144d67cb5dd5603XXX"

	mock := &mocks.MockClient{}

	client := inspect.Client{
		Client: mock,
	}

	stubWithFixture(t, mock, 200, "token.json")
	stubWithFixture(t, mock, 200, "manifest.json")
	stubWithFixture(t, mock, 404, "unknown-blob.json")

	_, err := client.Labels(repo, tag)
	assert.Error(t, err)
}

func stubWithFixture(t *testing.T, mock *mocks.MockClient, responseCode int, file string) {
	data, err := ioutil.ReadFile(path.Join("testdata", file))
	assert.NoError(t, err)

	// create a new reader with that JSON
	r := ioutil.NopCloser(bytes.NewReader(data))

	mock.StubResponse(responseCode, r)
}
