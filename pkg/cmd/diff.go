package cmd

import (
	"fmt"
	"github.com/garethjevans/inspect/pkg/registry"
	"sort"
	"strings"

	"github.com/garethjevans/inspect/pkg/util"

	"github.com/fatih/color"

	"github.com/garethjevans/inspect/pkg/inspect"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	colorError = color.New(color.FgRed).SprintFunc()
)

// DiffCmd a struct for the diff command.
type DiffCmd struct {
	BaseCmd
	Cmd    *cobra.Command
	Args   []string
	LabelLister registry.LabelLister
}

// NewDiffCmd creates a new diff command.
func NewDiffCmd() *cobra.Command {
	c := &DiffCmd{
		LabelLister: &registry.DefaultLabelLister{},
	}

	c.Log = c

	cmd := &cobra.Command{
		Use:     "diff <image> <image>",
		Short:   "Diff two docker images",
		Long:    "",
		Example: "inspect diff jenkinciinfra/terraform:1.0.0 jenkinsciinfra/terraform:1.1.0",
		Aliases: []string{"compare"},
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

// Run runs the command.
func (c *DiffCmd) Run() error {
	image1 := c.Args[0]
	image2 := c.Args[1]

	logrus.Debugf("comparing %s and %s", image1, image2)
	repo1, tag1, err := ParseRepo(image1)
	if err != nil {
		return err
	}

	repo2, tag2, err := ParseRepo(image2)
	if err != nil {
		return err
	}

	if repo1 != repo2 {
		return fmt.Errorf("images do not appear to be from the git repo 1=%s, 2=%s", repo1, repo2)
	}

	labels1, err := c.LabelLister.Labels(repo1, tag1)
	if err != nil {
		return err
	}

	labels2, err := c.LabelLister.Labels(repo2, tag2)
	if err != nil {
		return err
	}

	t := table.NewWriter()
	sb := strings.Builder{}
	t.SetOutputMirror(&sb)
	t.SetStyle(tableStyle)

	if headers {
		t.AppendHeader(table.Row{"Image", "1", "2"})
	}

	if writeSeparators {
		t.AppendSeparator()
	}

	t.AppendRow(table.Row{repo1, tag1, tag2})

	if writeSeparators {
		t.AppendSeparator()
	}

	keys := util.AllKeys(labels1, labels2)
	sort.Strings(keys)

	for _, k := range keys {
		if labels1[k] == labels2[k] {
			t.AppendRow(table.Row{k, labels1[k], labels2[k]})
		} else {
			t.AppendRow(table.Row{k, colorError(labels1[k]), colorError(labels2[k])})
		}
	}

	if writeSeparators {
		t.AppendSeparator()
	}
	t.AppendRow(table.Row{
		"GitHub URL",
		inspect.GitHubURL(labels1),
		inspect.GitHubURL(labels2),
	})

	if enableMarkdown {
		t.RenderMarkdown()
	} else {
		t.Render()
	}

	// write the table
	c.Log.Println(sb.String())

	// write the compare link
	c.Log.Println(fmt.Sprintf("%s/compare/%s..%s", inspect.BaseURL(labels1), inspect.Revision(labels1), inspect.Revision(labels2)))

	return nil
}
