package cmd

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// LabelsCmd struct for the labels command.
type LabelsCmd struct {
	BaseCmd
	Cmd  *cobra.Command
	Args []string

	IncludeGoVersion bool
}

// NewLabelsCmd struct for the labels command.
func NewLabelsCmd() *cobra.Command {
	c := &LabelsCmd{
		BaseCmd: NewBaseCmd(),
	}

	c.Log = c

	cmd := &cobra.Command{
		Use:     "labels",
		Short:   "Generates labels when creating an image",
		Long:    "",
		Example: "docker build $(inspect labels) ...",
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

	cmd.Flags().BoolVarP(&c.IncludeGoVersion, "include-go-version", "", false, "Attempt to include `go version` in the label set")

	return cmd
}

// Run runs the command.
func (c *LabelsCmd) Run() error {
	commands := []string{}

	gitCommitRev, err := c.GitCommitRev()
	if err != nil {
		return err
	}

	commands = append(commands, fmt.Sprintf("\"org.opencontainers.image.revision=%s\"", gitCommitRev))
	commands = append(commands, fmt.Sprintf("\"org.label-schema.vcs-ref=%s\"", gitCommitRev))

	gitScmURL, err := c.GitScmURL()
	if err != nil {
		return err
	}

	commands = append(commands, fmt.Sprintf("\"org.opencontainers.image.url=%s\"", gitScmURL))
	commands = append(commands, fmt.Sprintf("\"org.label-schema.url=%s\"", gitScmURL))

	buildDate, err := c.BuildDate()
	if err != nil {
		return err
	}

	commands = append(commands, fmt.Sprintf("\"org.opencontainers.image.created=%s\"", buildDate))
	commands = append(commands, fmt.Sprintf("\"org.label-schema.build-date=%s\"", buildDate))

	if c.IncludeGoVersion {
		goVersion, err := c.GoVersion()
		if err != nil {
			return err
		}
		commands = append(commands, fmt.Sprintf("\"inspect.tools.go.version=%s\"", goVersion))
	}

	gitTreeState, err := c.GitTreeState()
	if err != nil {
		return err
	}

	commands = append(commands, fmt.Sprintf("\"inspect.tree.state=%s\"", gitTreeState))

	c.Log.Println("--label " + strings.Join(commands, " --label "))

	return nil
}
