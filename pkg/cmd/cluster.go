package cmd

import (
	"errors"
	"strings"

	"github.com/garethjevans/inspect/pkg/registry"

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
	LabelLister registry.LabelLister

	Namespace string
}

// NewClusterCmd creates a new cluster command.
func NewClusterCmd() *cobra.Command {
	c := &ClusterCmd{
		BaseCmd: BaseCmd{
			CommandRunner: util.DefaultCommandRunner{},
		},
		LabelLister: &registry.DefaultLabelLister{},
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
		logrus.Debugf("checking %s", a)
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
