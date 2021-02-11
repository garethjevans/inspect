package cmd_test

import (
	"testing"

	"github.com/garethjevans/inspect/pkg/cmd/mock"

	"github.com/garethjevans/inspect/pkg/cmd"
	"github.com/stretchr/testify/assert"
)

func TestDiff(t *testing.T) {
	logger := &mock.LoggerMock{}
	//commandRunner := &mocks.MockCommandRunner{}

	c := cmd.DiffCmd{
		Log: logger,
		//BaseCmd: cmd.BaseCmd{
		//	CommandRunner: commandRunner,
		//},
	}

	c.Args = []string{"jenkinsciinfra/terraform:1.0.0", "jenkinsciinfra/terraform:1.1.0"}

	err := c.Run()
	assert.NoError(t, err)

	assert.Equal(t, 1, len(logger.Messages))
	assert.Equal(t, "https://github.com/jenkins-infra/docker-terraform/compare/ad902ec..441c261", logger.Messages[0])
}
