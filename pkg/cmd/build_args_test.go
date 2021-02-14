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

func TestBuildArgs(t *testing.T) {
	logger := &mock.LoggerMock{}
	commandRunner := &mocks.MockCommandRunner{}
	mocks.GetRunWithoutRetryFunc = func(c *util.Command) (string, error) {
		t.Log(c.String())
		if c.String() == gitRevParseShortHead {
			return gitRevParseShortHeadOutput, nil
		}
		if c.String() == gitConfigGetRemoteOriginURL {
			return gitConfigGetRemoteOriginURLOutput, nil
		}
		if c.String() == dateUtc {
			return dateUtcOutput, nil
		}
		if c.String() == gitStatusPorcelain {
			return gitStatusPercelainOutput, nil
		}
		return "", errors.New("unknown command")
	}

	c := cmd.BuildArgsCmd{
		BaseCmd: cmd.BaseCmd{
			Log:           logger,
			CommandRunner: commandRunner,
		},
	}

	err := c.Run()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(logger.Messages))
	assert.Equal(t, "--build-arg \"GIT_COMMIT_REV=sha123456\""+
		" --build-arg \"GIT_SCM_URL=https://github.com/org/repo.git\""+
		" --build-arg \"BUILD_DATE=2021-02-09T15:42:29\""+
		" --build-arg \"GIT_TREE_STATE=dirty\"", logger.Messages[0])
}

func TestBuildArgs_WithGoVersion(t *testing.T) {
	logger := &mock.LoggerMock{}
	commandRunner := &mocks.MockCommandRunner{}
	mocks.GetRunWithoutRetryFunc = func(c *util.Command) (string, error) {
		t.Log(c.String())
		if c.String() == gitRevParseShortHead {
			return gitRevParseShortHeadOutput, nil
		}
		if c.String() == gitConfigGetRemoteOriginURL {
			return gitConfigGetRemoteOriginURLOutput, nil
		}
		if c.String() == dateUtc {
			return dateUtcOutput, nil
		}
		if c.String() == goVersion {
			return goVersionOutput, nil
		}
		if c.String() == gitStatusPorcelain {
			return gitStatusPercelainOutput, nil
		}
		return "", errors.New("unknown command")
	}

	c := cmd.BuildArgsCmd{
		BaseCmd: cmd.BaseCmd{
			Log:           logger,
			CommandRunner: commandRunner,
		},
		IncludeGoVersion: true,
	}

	err := c.Run()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(logger.Messages))
	assert.Equal(t, "--build-arg \"GIT_COMMIT_REV=sha123456\""+
		" --build-arg \"GIT_SCM_URL=https://github.com/org/repo.git\""+
		" --build-arg \"BUILD_DATE=2021-02-09T15:42:29\""+
		" --build-arg \"GO_VERSION=1.15.6\""+
		" --build-arg \"GIT_TREE_STATE=dirty\"", logger.Messages[0])
}
