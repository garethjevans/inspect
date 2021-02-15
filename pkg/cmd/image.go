package cmd

import (
	"errors"
	"strings"

	"github.com/garethjevans/inspect/pkg/registry"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ImageCmd struct for the image command.
type ImageCmd struct {
	BaseCmd
	Cmd         *cobra.Command
	Args        []string
	LabelLister registry.LabelLister
}

// NewImageCmd creates a new ImageCmd.
func NewImageCmd() *cobra.Command {
	c := &ImageCmd{
		LabelLister: &registry.DefaultLabelLister{},
	}
	c.Log = c
	cmd := &cobra.Command{
		Use:     "image <name>...",
		Short:   "Inspect the docker container",
		Long:    "",
		Example: "",
		Aliases: []string{"i", "in", "ins"},
		Run: func(cmd *cobra.Command, args []string) {
			c.Cmd = cmd
			c.Args = args
			err := c.Run()
			if err != nil {
				logrus.Fatalf("unable to run command: %s", err)
			}
		},
		Args: cobra.MinimumNArgs(1),
	}

	return cmd
}

// Run runs the command.
func (c *ImageCmd) Run() error {
	for _, a := range c.Args {
		repo, tag := ParseRepo(a)

		if repo == "" {
			return errors.New("no repository has been configured")
		}

		if tag == "" {
			return errors.New("no tag has been configured")
		}

		labels, err := c.LabelLister.Labels(repo, tag)
		if err != nil {
			return err
		}

		if len(labels) == 0 {
			c.Log.Println("No labels found for " + a)
		} else {
			t := table.NewWriter()
			sb := strings.Builder{}
			t.SetOutputMirror(&sb)

			writeTableForImage(labels, t, repo, tag)

			if enableMarkdown {
				t.RenderMarkdown()
			} else {
				t.Render()
			}

			// write the table
			c.Log.Println(sb.String())
		}
	}
	return nil
}
