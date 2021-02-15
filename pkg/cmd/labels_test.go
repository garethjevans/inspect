package cmd_test

import (
	"errors"
	"testing"

	mocks2 "github.com/garethjevans/inspect/pkg/mocks"

	"github.com/garethjevans/inspect/pkg/cmd"
	"github.com/garethjevans/inspect/pkg/util"
	"github.com/stretchr/testify/assert"
)

const (
	gitRevParseShortHead              = "git rev-parse --short HEAD"
	gitConfigGetRemoteOriginURL       = "git config --get remote.origin.url"
	dateUtc                           = "date --utc +%Y-%m-%dT%H:%M:%S"
	goVersion                         = "go version"
	gitStatusPorcelain                = "git status --porcelain"
	gitConfigGetRemoteOriginURLOutput = "https://github.com/org/repo.git"
	gitRevParseShortHeadOutput        = "sha123456"
	dateUtcOutput                     = "2021-02-09T15:42:29"
	goVersionOutput                   = "go version go1.15.6 darwin/amd64"
	gitStatusPercelainOutput          = "?? ll"
)

func TestLabelsCmd(t *testing.T) {
	logger := &mocks2.LoggerMock{}
	commandRunner := &mocks2.MockCommandRunner{}
	mocks2.GetRunWithoutRetryFunc = func(c *util.Command) (string, error) {
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
		" --label \"inspect.tree.state=dirty\"", logger.Messages[0])
}

func TestLabelsCmd_WithGoVersion(t *testing.T) {
	logger := &mocks2.LoggerMock{}
	commandRunner := &mocks2.MockCommandRunner{}
	mocks2.GetRunWithoutRetryFunc = func(c *util.Command) (string, error) {
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

	c := cmd.LabelsCmd{
		BaseCmd: cmd.BaseCmd{
			Log:           logger,
			CommandRunner: commandRunner,
		},
		IncludeGoVersion: true,
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
		" --label \"inspect.tools.go.version=1.15.6\""+
		" --label \"inspect.tree.state=dirty\"", logger.Messages[0])
}
