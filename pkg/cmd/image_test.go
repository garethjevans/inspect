package cmd_test

import (
	"testing"

	"github.com/garethjevans/inspect/pkg/inspect"
	"github.com/garethjevans/inspect/pkg/inspect/mocks"

	"github.com/garethjevans/inspect/pkg/cmd"
	"github.com/stretchr/testify/assert"
)

func TestImage(t *testing.T) {
	mock := &mocks.MockClient{}

	c := cmd.ImageCmd{
		Client: inspect.Client{
			Client: mock,
		},
	}

	stubWithFixture(t, mock, "token.json")
	stubWithFixture(t, mock, "manifests.1.0.0.json")
	stubWithFixture(t, mock, "blobs.1.0.0.json")

	c.Args = []string{"jenkinsciinfra/terraform:1.0.0"}

	err := c.Run()
	assert.NoError(t, err)
}
