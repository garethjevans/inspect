package cmd

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/garethjevans/inspect/pkg/util"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type BuildArgsCmd struct {
	Cmd           *cobra.Command
	Args          []string
	CommandRunner util.CommandRunner
	Log           Logs
}

func NewBuildArgsCmd() *cobra.Command {
	c := &BuildArgsCmd{
		CommandRunner: util.DefaultCommandRunner{},
	}

	c.Log = c

	cmd := &cobra.Command{
		Use:     "build-args",
		Short:   "Generates build args when creating an image",
		Long:    "",
		Example: "",
		Aliases: []string{"args"},
		Run: func(cmd *cobra.Command, args []string) {
			c.Cmd = cmd
			c.Args = args
			err := c.Run()
			if err != nil {
				logrus.Fatalf("unable to run command: %s", err)
			}
		},
		Args: cobra.NoArgs,
	}

	return cmd
}

func (c *BuildArgsCmd) Run() error {
	gitCommitRevCommand := util.Command{
		Name: "git",
		Args: []string{"log", "-n", "1", "--pretty=format:%h"},
	}

	out, err := c.CommandRunner.RunWithoutRetry(&gitCommitRevCommand)
	if err != nil {
		return err
	}

	c.Log.Println(fmt.Sprintf("GIT_COMMIT_REV=%s", out))

	gitScmURLCommand := util.Command{
		Name: "git",
		Args: []string{"config", "--get", "remote.origin.url"},
	}

	out, err = c.CommandRunner.RunWithoutRetry(&gitScmURLCommand)
	if err != nil {
		return err
	}

	out = strings.ReplaceAll(out, "git@github.com:", "https://github.com/")

	c.Log.Println(fmt.Sprintf("GIT_SCM_URL=%s", out))

	buildDate := util.Command{
		Name: "date",
		Args: []string{"--utc", "+%Y-%m-%dT%H:%M:%S"},
	}

	out, err = c.CommandRunner.RunWithoutRetry(&buildDate)
	if err != nil {
		return err
	}

	c.Log.Println(fmt.Sprintf("BUILD_DATE=%s", out))

	goVersionCommand := util.Command{
		Name: "go",
		Args: []string{"version"},
	}

	out, err = c.CommandRunner.RunWithoutRetry(&goVersionCommand)
	if err != nil {
		return err
	}

	re := regexp.MustCompile(`\d+(\.\d+)+`)
	goVersion := re.FindString(out)
	c.Log.Println(fmt.Sprintf("GO_VERSION=%s", goVersion))

	gitStatusCommand := util.Command{
		Name: "git",
		Args: []string{"status", "--porcelain"},
	}

	out, err = c.CommandRunner.RunWithoutRetry(&gitStatusCommand)
	if err != nil {
		return err
	}

	var gitTreeState string
	if out == "" {
		gitTreeState = "clean"
	} else {
		gitTreeState = "dirty"
	}

	c.Log.Println(fmt.Sprintf("GIT_TREE_STATE=%s", gitTreeState))

	//--label "org.opencontainers.image.source=$(GIT_SCM_URL)" \
	//--label "org.label-schema.vcs-url=$(GIT_SCM_URL)" \
	//--label "org.opencontainers.image.url=$(SCM_URI)" \
	//--label "org.label-schema.url=$(SCM_URI)" \
	//--label "org.opencontainers.image.revision=$(GIT_COMMIT_REV)" \
	//--label "org.label-schema.vcs-ref=$(GIT_COMMIT_REV)" \
	//--label "org.opencontainers.image.created=$(BUILD_DATE)" \
	//--label "org.label-schema.build-date=$(BUILD_DATE)" \

	return nil
}

type Logs interface {
	Println(message string)
}

func (c *BuildArgsCmd) Println(message string) {
	fmt.Println(message)
}
