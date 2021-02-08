package cmd

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/garethjevans/inspect/pkg/inspect"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type InspectCmd struct {
	Cmd  *cobra.Command
	Args []string

	Repository string
	Tag        string
}

func NewInspectCmd() *cobra.Command {
	c := &InspectCmd{}
	cmd := &cobra.Command{
		Use:     "inspect",
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
		Args: cobra.NoArgs,
	}

	cmd.Flags().StringVarP(&c.Repository, "repository", "r", "",
		"Repository to query")
	cmd.Flags().StringVarP(&c.Tag, "tag", "t", "",
		"Tag to query")

	return cmd
}

func (c *InspectCmd) Run() error {
	client := inspect.Client{
		Client: http.Client{},
	}

	labels, err := client.Labels(c.Repository, c.Tag)
	if err != nil {
		return err
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Label", "Value"})
	t.AppendSeparator()

	for k, v := range labels {
		t.AppendRow(table.Row{k, v})
	}

	gitURL := labels["org.opencontainers.image.source"]
	if strings.HasSuffix(gitURL, ".git") {
		gitURL = strings.TrimSuffix(gitURL, ".git")
	}

	possibleURL := fmt.Sprintf("%s/tree/%s", gitURL, labels["org.opencontainers.image.revision"])
	t.AppendSeparator()
	t.AppendRow(table.Row{"GitHub URL", possibleURL})

	t.Render()

	return nil
}
