package cmd

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/garethjevans/inspect/pkg/util"
)

// BaseCmd a common type to contain all the git helper methods.
type BaseCmd struct {
	CommandRunner util.CommandRunner
	Log           Logger
}

// GitCommitRev helper to return the git commit rev.
func (b *BaseCmd) GitCommitRev() (string, error) {
	gitCommitRevCommand := util.Command{
		Name: "git",
		Args: []string{"rev-parse", "--short", "HEAD"},
	}

	out, err := b.CommandRunner.RunWithoutRetry(&gitCommitRevCommand)
	if err != nil {
		return "", err
	}

	return out, nil
}

// GitScmURL helper to return the git scm url.
func (b *BaseCmd) GitScmURL() (string, error) {
	gitScmURLCommand := util.Command{
		Name: "git",
		Args: []string{"config", "--get", "remote.origin.url"},
	}

	out, err := b.CommandRunner.RunWithoutRetry(&gitScmURLCommand)
	if err != nil {
		return "", err
	}

	out = strings.ReplaceAll(out, "git@github.com:", "https://github.com/")
	return out, nil
}

// BuildDate helper to return the current build date.
func (b *BaseCmd) BuildDate() (string, error) {
	buildDate := util.Command{
		Name: "date",
		Args: []string{"--utc", "+%Y-%m-%dT%H:%M:%S"},
	}

	out, err := b.CommandRunner.RunWithoutRetry(&buildDate)
	if err != nil {
		return "", err
	}

	return out, nil
}

// GoVersion helper to return the current go version.
func (b *BaseCmd) GoVersion() (string, error) {
	goVersionCommand := util.Command{
		Name: "go",
		Args: []string{"version"},
	}

	out, err := b.CommandRunner.RunWithoutRetry(&goVersionCommand)
	if err != nil {
		return "", err
	}

	re := regexp.MustCompile(`\d+(\.\d+)+`)
	goVersion := re.FindString(out)
	return goVersion, nil
}

// GitTreeState helper to return the current tree state (clean|dirty).
func (b *BaseCmd) GitTreeState() (string, error) {
	gitStatusCommand := util.Command{
		Name: "git",
		Args: []string{"status", "--porcelain"},
	}

	out, err := b.CommandRunner.RunWithoutRetry(&gitStatusCommand)
	if err != nil {
		return "", err
	}

	var gitTreeState string
	if out == "" {
		gitTreeState = "clean"
	} else {
		gitTreeState = "dirty"
	}

	return gitTreeState, nil
}

// Println a helper to allow this to be mocked out.
func (b *BaseCmd) Println(message string) {
	fmt.Println(message)
}
