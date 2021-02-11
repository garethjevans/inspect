package cmd

import (
	"fmt"

	"github.com/garethjevans/inspect/pkg/util"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type ClusterCmd struct {
	BaseCmd
	Cmd  *cobra.Command
	Args []string
	Log  Logs
}

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

func (c *ClusterCmd) Run() error {
	// connect with local kubeconfig

	//

	return nil
}

func (c *ClusterCmd) Println(message string) {
	fmt.Println(message)
}
