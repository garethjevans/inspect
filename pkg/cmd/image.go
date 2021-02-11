package cmd

import (
	"errors"
	"net/http"
	"os"

	"github.com/garethjevans/inspect/pkg/inspect"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ImageCmd struct for the image command.
type ImageCmd struct {
	Cmd    *cobra.Command
	Args   []string
	Client inspect.Client
}

// NewImageCmd creates a new ImageCmd.
func NewImageCmd() *cobra.Command {
	c := &ImageCmd{
		Client: inspect.Client{
			Client: &http.Client{},
		},
	}
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
		Args: cobra.MaximumNArgs(1),
	}

	return cmd
}

// Run runs the command.
func (c *ImageCmd) Run() error {
	for _, a := range c.Args {
		repo, tag, err := ParseRepo(a)
		if err != nil {
			return err
		}

		if repo == "" {
			return errors.New("no repository has been configured")
		}

		if tag == "" {
			return errors.New("no tag has been configured")
		}

		labels, err := c.Client.Labels(repo, tag)
		if err != nil {
			return err
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.SetStyle(TableStyle)

		if Headers {
			t.AppendHeader(table.Row{"Label", "Value"})
		}

		if WriteSeparators {
			t.AppendSeparator()
		}

		for k, v := range labels {
			t.AppendRow(table.Row{k, v})
		}

		if WriteSeparators {
			t.AppendSeparator()
		}
		t.AppendRow(table.Row{"GitHub URL", inspect.GitHubURL(labels)})

		if EnableMarkdown {
			t.RenderMarkdown()
		} else {
			t.Render()
		}
	}
	return nil
}
