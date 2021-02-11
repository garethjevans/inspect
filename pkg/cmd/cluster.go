package cmd

import (
	"github.com/garethjevans/inspect/pkg/util"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ClusterCmd struct for the cluster command.
type ClusterCmd struct {
	BaseCmd
	Cmd  *cobra.Command
	Args []string
}

// NewClusterCmd creates a new cluster command.
func NewClusterCmd() *cobra.Command {
	c := &ClusterCmd{
		BaseCmd: BaseCmd{
			CommandRunner: util.DefaultCommandRunner{},
		},
	}

	c.Log = c

	cmd := &cobra.Command{
		Use:     "cluster",
		Short:   "Inspect all containers running in a cluster",
		Long:    "",
		Example: "",
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

	return cmd
}

// Run runs the command.
func (c *ClusterCmd) Run() error {
	// connect with local kubeconfig

	// TODO

	// get a list of all pods

	// get the images for all containers in the pod.

	// extract the labels for each

	return nil
}
