package cmd_test

import (
	"errors"
	"testing"

	"github.com/garethjevans/inspect/pkg/cmd/mock"

	"github.com/garethjevans/inspect/pkg/util"
	"github.com/garethjevans/inspect/pkg/util/mocks"

	"github.com/garethjevans/inspect/pkg/cmd"
	"github.com/stretchr/testify/assert"
)

func TestLabelsCmd(t *testing.T) {
	logger := &mock.LoggerMock{}
	commandRunner := &mocks.MockCommandRunner{}
	mocks.GetRunWithoutRetryFunc = func(c *util.Command) (string, error) {
		t.Log(c.String())
		if c.String() == "git rev-parse --short HEAD" {
			return "sha123456", nil
		}
		if c.String() == "git config --get remote.origin.url" {
			return "https://github.com/org/repo.git", nil
		}
		if c.String() == "date --utc +%Y-%m-%dT%H:%M:%S" {
			return "2021-02-09T15:42:29", nil
		}
		if c.String() == "go version" {
			return "go version go1.15.6 darwin/amd64", nil
		}
		if c.String() == "git status --porcelain" {
			return "?? ll", nil
		}
		return "", errors.New("unknown command")
	}

	c := cmd.LabelsCmd{
		BaseCmd: cmd.BaseCmd{
			Log:           logger,
			CommandRunner: commandRunner,
		},
	}

	err := c.Run()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(logger.Messages))
	assert.Equal(t, "--label \"org.opencontainers.image.revision=sha123456\""+
		" --label \"org.label-schema.vcs-ref=sha123456\""+
		" --label \"org.opencontainers.image.url=https://github.com/org/repo.git\""+
		" --label \"org.label-schema.url=https://github.com/org/repo.git\""+
		" --label \"org.opencontainers.image.created=2021-02-09T15:42:29\""+
		" --label \"org.label-schema.build-date=2021-02-09T15:42:29\""+
		" --label \"io.jenkins-infra.go.version=1.15.6\""+
		" --label \"io.jenkins-infra.tree.state=dirty\"", logger.Messages[0])
}
