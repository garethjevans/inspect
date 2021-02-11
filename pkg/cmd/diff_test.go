package cmd_test

import (
	"bytes"
	"io/ioutil"
	"path"
	"testing"

	"github.com/garethjevans/inspect/pkg/inspect"
	"github.com/garethjevans/inspect/pkg/inspect/mocks"

	"github.com/garethjevans/inspect/pkg/cmd/mock"

	"github.com/garethjevans/inspect/pkg/cmd"
	"github.com/stretchr/testify/assert"
)

func TestDiff(t *testing.T) {
	logger := &mock.LoggerMock{}
	mock := &mocks.MockClient{}

	c := cmd.DiffCmd{
		Log: logger,
		Client: inspect.Client{
			Client: mock,
		},
	}

	stubWithFixture(t, mock, "token.json")
	stubWithFixture(t, mock, "manifests.1.0.0.json")
	stubWithFixture(t, mock, "blobs.1.0.0.json")

	stubWithFixture(t, mock, "token.json")
	stubWithFixture(t, mock, "manifests.1.1.0.json")
	stubWithFixture(t, mock, "blobs.1.1.0.json")

	c.Args = []string{"jenkinsciinfra/terraform:1.0.0", "jenkinsciinfra/terraform:1.1.0"}

	err := c.Run()
	assert.NoError(t, err)

	assert.Equal(t, 1, len(logger.Messages))
	assert.Equal(t, "https://github.com/jenkins-infra/docker-terraform/compare/ad902ec..441c261", logger.Messages[0])
}

func stubWithFixture(t *testing.T, mock *mocks.MockClient, file string) {
	data, err := ioutil.ReadFile(path.Join("testdata", file))
	assert.NoError(t, err)

	// create a new reader with that JSON
	r := ioutil.NopCloser(bytes.NewReader(data))

	mock.StubResponse(200, r)
}
