package cmd_test

import (
	"errors"
	"testing"

	"github.com/garethjevans/inspect/pkg/util"
	"github.com/garethjevans/inspect/pkg/util/mocks"

	"github.com/garethjevans/inspect/pkg/cmd"
	"github.com/stretchr/testify/assert"
)

type LoggerMock struct {
	messages []string
}

func (l *LoggerMock) Println(message string) {
	l.messages = append(l.messages, message)
}

func TestBuildArgs(t *testing.T) {
	logger := &LoggerMock{}
	commandRunner := &mocks.MockCommandRunner{}
	mocks.GetRunWithoutRetryFunc = func(c *util.Command) (string, error) {
		t.Log(c.String())
		if c.String() == "git log -n 1 --pretty=format:%h" {
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

	c := cmd.BuildArgsCmd{
		Log:           logger,
		CommandRunner: commandRunner,
	}

	err := c.Run()
	assert.NoError(t, err)
	assert.Equal(t, 5, len(logger.messages))
	assert.Equal(t, "GIT_COMMIT_REV=sha123456", logger.messages[0])
	assert.Equal(t, "GIT_SCM_URL=https://github.com/org/repo.git", logger.messages[1])
	assert.Equal(t, "BUILD_DATE=2021-02-09T15:42:29", logger.messages[2])
	assert.Equal(t, "GO_VERSION=1.15.6", logger.messages[3])
	assert.Equal(t, "GIT_TREE_STATE=dirty", logger.messages[4])
}
