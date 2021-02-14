package cmd

import (
	"errors"
	"net/http"
	"sort"
	"strings"

	"github.com/garethjevans/inspect/pkg/util"

	"github.com/garethjevans/inspect/pkg/inspect"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ImageCmd struct for the image command.
type ImageCmd struct {
	BaseCmd
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

		if len(labels) == 0 {
			c.Log.Println("No labels found for " + a)
		} else {
			t := table.NewWriter()

			sb := strings.Builder{}
			t.SetOutputMirror(&sb)
			t.SetStyle(tableStyle)

			if headers {
				t.AppendHeader(table.Row{"Label", "Value"})
			}

			if writeSeparators {
				t.AppendSeparator()
			}

			keys := util.AllKeys(labels)
			sort.Strings(keys)

			for _, k := range keys {
				t.AppendRow(table.Row{k, labels[k]})
			}

			if writeSeparators {
				t.AppendSeparator()
			}

			t.AppendRow(table.Row{"GitHub URL", inspect.GitHubURL(labels)})

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
