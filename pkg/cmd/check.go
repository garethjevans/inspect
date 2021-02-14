package cmd

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/garethjevans/inspect/pkg/inspect"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// CheckCmd struct for the check command.
type CheckCmd struct {
	BaseCmd
	Cmd                  *cobra.Command
	Args                 []string
	Client               inspect.Client
	FailOnRecommendation bool
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
		Use:     "check <image>...",
		Short:   "Check the image for recommended labels",
		Long:    "Check the images for recommended labels, provides a tabular output with recommendations if a particular label does not exist.",
		Example: "inspect check alpine:3.13.0",
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

	cmd.Flags().BoolVarP(&c.FailOnRecommendation, "fail-on-recommendations", "f", false, "Should exit 1 if there are recommendations")
	return cmd
}

// Run runs the command.
func (c *CheckCmd) Run() error {
	r := false
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

		r = recommendationRow("org.opencontainers.image.created", "date --utc +%Y-%m-%dT%H:%M:%S", labels, t, r)
		r = recommendationRow("org.opencontainers.image.revision", "git rev-parse --short HEAD", labels, t, r)
		r = recommendationRow("org.opencontainers.image.source", "git config --get remote.origin.url", labels, t, r)
		r = recommendationRow("org.opencontainers.image.url", "git config --get remote.origin.url", labels, t, r)
		r = recommendationRow("org.label-schema.build-date", "date --utc +%Y-%m-%dT%H:%M:%S", labels, t, r)
		r = recommendationRow("org.label-schema.vcs-ref", "git rev-parse --short HEAD", labels, t, r)
		r = recommendationRow("org.label-schema.vcs-url", "git config --get remote.origin.url", labels, t, r)
		r = recommendationRow("org.label-schema.url", "git config --get remote.origin.url", labels, t, r)

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

	if c.FailOnRecommendation && r {
		os.Exit(1)
	}

	return nil
}

func recommendationRow(name string, recommendation string, labels map[string]string, t table.Writer, recommendations bool) bool {
	created := labels[name]
	if created == "" {
		t.AppendRow(table.Row{name, "Missing", recommendation})
		recommendations = true
	} else {
		t.AppendRow(table.Row{name, "OK", ""})
	}
	return recommendations
}
