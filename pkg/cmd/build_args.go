package cmd

import (
	"fmt"
	"strings"

	"github.com/garethjevans/inspect/pkg/util"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type BuildArgsCmd struct {
	BaseCmd
	Cmd  *cobra.Command
	Args []string
	Log  Logs
}

func NewBuildArgsCmd() *cobra.Command {
	c := &BuildArgsCmd{
		BaseCmd: BaseCmd{
			CommandRunner: util.DefaultCommandRunner{},
		},
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
	commands := []string{}

	gitCommitRev, err := c.GitCommitRev()
	if err != nil {
		return err
	}

	commands = append(commands, fmt.Sprintf("\"GIT_COMMIT_REV=%s\"", gitCommitRev))

	gitScmURL, err := c.GitScmURL()
	if err != nil {
		return err
	}

	commands = append(commands, fmt.Sprintf("\"GIT_SCM_URL=%s\"", gitScmURL))

	buildDate, err := c.BuildDate()
	if err != nil {
		return err
	}

	commands = append(commands, fmt.Sprintf("\"BUILD_DATE=%s\"", buildDate))

	goVersion, err := c.GoVersion()
	if err != nil {
		return err
	}

	commands = append(commands, fmt.Sprintf("\"GO_VERSION=%s\"", goVersion))

	gitTreeState, err := c.GitTreeState()
	if err != nil {
		return err
	}

	commands = append(commands, fmt.Sprintf("\"GIT_TREE_STATE=%s\"", gitTreeState))

	c.Log.Println("--build-arg " + strings.Join(commands, " --build-arg "))

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

func (c *BuildArgsCmd) Println(message string) {
	fmt.Println(message)
}
