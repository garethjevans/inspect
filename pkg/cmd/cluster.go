package cmd

import (
	"errors"
	"net/http"
	"sort"
	"strings"

	"github.com/garethjevans/inspect/pkg/inspect"
	"github.com/garethjevans/inspect/pkg/kube"
	"github.com/garethjevans/inspect/pkg/util"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ClusterCmd struct for the cluster command.
type ClusterCmd struct {
	BaseCmd
	Cmd         *cobra.Command
	Args        []string
	ImageLister kube.ImageLister
	Client      inspect.Client

	Namespace string
}

// NewClusterCmd creates a new cluster command.
func NewClusterCmd() *cobra.Command {
	c := &ClusterCmd{
		BaseCmd: BaseCmd{
			CommandRunner: util.DefaultCommandRunner{},
		},
		Client: inspect.Client{
			Client: &http.Client{},
		},
		ImageLister: &kube.Kuber{},
	}

	c.Log = c

	cmd := &cobra.Command{
		Use:     "cluster",
		Short:   "Inspect all containers running in a cluster",
		Long:    "",
		Example: "inspect cluster --namespace <mynamespace>",
		//Aliases: []string{""},
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

	cmd.Flags().StringVarP(&c.Namespace, "namespace", "n", "", "Namespace to filter on")

	return cmd
}

// Run runs the command.
func (c *ClusterCmd) Run() error {
	images, err := c.ImageLister.GetImagesForNamespace(c.Namespace)
	if err != nil {
		return err
	}

	// extract the labels for each
	for _, a := range images {
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

			t.AppendRow(table.Row{repo, tag})

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
