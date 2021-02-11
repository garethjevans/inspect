package cmd

import (
	"errors"
	"net/http"
	"strings"

	"github.com/garethjevans/inspect/pkg/inspect"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// CheckCmd struct for the check command.
type CheckCmd struct {
	BaseCmd
	Cmd    *cobra.Command
	Args   []string
	Client inspect.Client
}

// NewCheckCmd creates a new CheckCmd.
func NewCheckCmd() *cobra.Command {
	c := &CheckCmd{
		Client: inspect.Client{
			Client: &http.Client{},
		},
	}
	c.Log = c
	cmd := &cobra.Command{
		Use:     "check <name>...",
		Short:   "Check the docker container for recommended labels",
		Long:    "",
		Example: "",
		Aliases: []string{"validate"},
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
func (c *CheckCmd) Run() error {
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

		sb := strings.Builder{}
		t.SetOutputMirror(&sb)
		t.SetStyle(tableStyle)

		if headers {
			t.AppendHeader(table.Row{"Label", "OK", "Recommendation"})
		}

		if writeSeparators {
			t.AppendSeparator()
		}

		created := labels["org.opencontainers.image.created"]
		if created == "" {
			t.AppendRow(table.Row{"org.opencontainers.image.created", "Missing", "date --utc +%Y-%m-%dT%H:%M:%S"})
		} else {
			t.AppendRow(table.Row{"org.opencontainers.image.created", "OK", ""})
		}

		revision := labels["org.opencontainers.image.revision"]
		if revision == "" {
			t.AppendRow(table.Row{"org.opencontainers.image.revision", "Missing", "git log -n 1 --pretty=format:%h"})
		} else {
			t.AppendRow(table.Row{"org.opencontainers.image.revision", "OK", ""})
		}

		source := labels["org.opencontainers.image.source"]
		if source == "" {
			t.AppendRow(table.Row{"org.opencontainers.image.source", "Missing", "git config --get remote.origin.url"})
		} else {
			t.AppendRow(table.Row{"org.opencontainers.image.source", "OK", ""})
		}

		url := labels["org.opencontainers.image.url"]
		if url == "" {
			t.AppendRow(table.Row{"org.opencontainers.image.url", "Missing", "git config --get remote.origin.url"})
		} else {
			t.AppendRow(table.Row{"org.opencontainers.image.url", "OK", ""})
		}

		if writeSeparators {
			t.AppendSeparator()
		}

		if enableMarkdown {
			t.RenderMarkdown()
		} else {
			t.Render()
		}

		// write the table
		c.Log.Println(sb.String())
	}

	return nil
}
