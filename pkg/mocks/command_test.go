package mocks_test

import (
	"errors"
	"testing"

	"github.com/garethjevans/inspect/pkg/mocks"
	"github.com/garethjevans/inspect/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestMockCommandRunner_RunWithoutRetry(t *testing.T) {
	runner := mocks.MockCommandRunner{}
	runner.RunWithoutRetryFunc = func(c *util.Command) (string, error) {
		if c.String() == "dummy one two three" {
			return "expected", nil
		}
		return "", errors.New("unexpected error")
	}

	cmd := util.Command{
		Name: "dummy",
		Args: []string{"one", "two", "three"},
	}

	output, err := runner.RunWithoutRetry(&cmd)
	assert.NoError(t, err)
	assert.Equal(t, output, "expected")

	cmd2 := util.Command{
		Name: "git",
		Args: []string{"one", "two", "three"},
	}

	output, err = runner.RunWithoutRetry(&cmd2)
	assert.Error(t, err)
	assert.Equal(t, output, "")
}
