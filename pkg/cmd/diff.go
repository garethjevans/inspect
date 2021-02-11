package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/fatih/color"

	"github.com/garethjevans/inspect/pkg/inspect"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	colorError = color.New(color.FgRed).SprintFunc()
)

type DiffCmd struct {
	Cmd    *cobra.Command
	Args   []string
	Log    Logs
	Client inspect.Client
}

func NewDiffCmd() *cobra.Command {
	c := &DiffCmd{
		Client: inspect.Client{
			Client: &http.Client{},
		},
	}

	c.Log = c

	cmd := &cobra.Command{
		Use:     "diff <image> <image>",
		Short:   "Diff two docker images",
		Long:    "",
		Example: "",
		//Aliases: []string{"i", "in", "ins"},
		Run: func(cmd *cobra.Command, args []string) {
			c.Cmd = cmd
			c.Args = args
			err := c.Run()
			if err != nil {
				logrus.Fatalf("unable to run command: %s", err)
			}
		},
		Args: cobra.ExactArgs(2),
	}

	return cmd
}

func (c *DiffCmd) Run() error {
	image1 := c.Args[0]
	image2 := c.Args[1]

	logrus.Debugf("comparing %s and %s", image1, image2)
	repo1, tag1 := ParseRepo(image1)
	repo2, tag2 := ParseRepo(image2)

	if repo1 != repo2 {
		return fmt.Errorf("images do not appear to be from the git repo 1=%s, 2=%s", repo1, repo2)
	}

	labels1, err := c.Client.Labels(repo1, tag1)
	if err != nil {
		return err
	}

	labels2, err := c.Client.Labels(repo2, tag2)
	if err != nil {
		return err
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Image", "1", "2"})
	t.AppendSeparator()
	t.AppendRow(table.Row{repo1, tag1, tag2})
	t.AppendSeparator()

	keys := AllKeys(labels1, labels2)

	for _, k := range keys {
		if labels1[k] == labels2[k] {
			t.AppendRow(table.Row{k, labels1[k], labels2[k]})
		} else {
			t.AppendRow(table.Row{k, colorError(labels1[k]), colorError(labels2[k])})
		}
	}

	t.AppendSeparator()
	t.AppendRow(table.Row{
		"GitHub URL",
		inspect.GitHubURL(labels1),
		inspect.GitHubURL(labels2),
	})

	t.Render()

	c.Log.Println(fmt.Sprintf("%s/compare/%s..%s", inspect.BaseURL(labels1), inspect.Revision(labels1), inspect.Revision(labels2)))

	return nil
}

// Println a helper to allow this to be mocked out.
func (c *DiffCmd) Println(message string) {
	fmt.Println(message)
}
