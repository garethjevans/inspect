package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/garethjevans/inspect/pkg/inspect"

	"github.com/garethjevans/inspect/pkg/kube"
	"github.com/garethjevans/inspect/pkg/registry"
	"github.com/garethjevans/inspect/pkg/util"
	"github.com/jedib0t/go-pretty/v6/table"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type comparison struct {
	image string
	tag1  string
	tag2  string
}

// DiffNamespaceCmd a struct for the diff namespace command.
type DiffNamespaceCmd struct {
	BaseCmd
	Cmd         *cobra.Command
	Args        []string
	LabelLister registry.LabelLister
	ImageLister kube.ImageLister
}

// NewDiffNamespaceCmd creates a new diff namespace command.
func NewDiffNamespaceCmd() *cobra.Command {
	c := &DiffNamespaceCmd{
		LabelLister: &registry.DefaultLabelLister{},
		ImageLister: &kube.Kuber{},
	}

	c.Log = c

	cmd := &cobra.Command{
		Use:     "diff-namespace <namespace1> <namespace2>",
		Short:   "Diff two namespaces",
		Long:    "",
		Example: "inspect diff-namespace <namespace1> <namespace2>",
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
func (c *DiffNamespaceCmd) Run() error {
	namespace1 := c.Args[0]
	namespace2 := c.Args[1]

	logrus.Debugf("comparing namespace %s and %s", namespace1, namespace2)

	imagesInNamespace1, err := c.ImageLister.GetImagesForNamespace(namespace1)
	if err != nil {
		return err
	}

	logrus.Debugf("found images %s in namespace %s", imagesInNamespace1, namespace1)

	imagesInNamespace2, err := c.ImageLister.GetImagesForNamespace(namespace2)
	if err != nil {
		return err
	}

	logrus.Debugf("found images %s in namespace %s", imagesInNamespace2, namespace2)

	// get all repository names
	imageNamesInNamespace1 := names(imagesInNamespace1)
	logrus.Debugf("filtered list down to %s", imageNamesInNamespace1)
	imageNamesInNamespace2 := names(imagesInNamespace2)
	logrus.Debugf("filtered list down to %s", imageNamesInNamespace2)

	// get a unique list of image names
	allImageNames := []string{}
	allImageNames = append(allImageNames, imageNamesInNamespace1...)
	allImageNames = append(allImageNames, imageNamesInNamespace2...)

	// sort
	unqiueImageNames := util.Unqiue(allImageNames)
	sort.Strings(unqiueImageNames)

	t := table.NewWriter()
	sb := strings.Builder{}
	t.SetOutputMirror(&sb)

	t.AppendRow(table.Row{"", namespace1, namespace2})
	t.AppendSeparator()

	toCompare := []comparison{}

	// loop through
	for _, i := range unqiueImageNames {
		versionInNamespace1 := locateVersion(imagesInNamespace1, i)
		versionInNamespace2 := locateVersion(imagesInNamespace2, i)
		t.AppendRow(table.Row{i, versionInNamespace1, versionInNamespace2})

		if versionInNamespace1 != "" && versionInNamespace2 != "" && versionInNamespace1 != versionInNamespace2 {
			logrus.Debugf("need to run comparison between %s:%s and %s:%s", i, versionInNamespace1, i, versionInNamespace2)
			toCompare = append(toCompare, comparison{image: i, tag1: versionInNamespace1, tag2: versionInNamespace2})
		}
	}

	t.Render()
	c.Log.Println(sb.String())

	for _, compare := range toCompare {
		t := table.NewWriter()
		sb := strings.Builder{}
		t.SetOutputMirror(&sb)

		labels1, err := c.LabelLister.Labels(compare.image, compare.tag1)
		if err != nil {
			return err
		}

		labels2, err := c.LabelLister.Labels(compare.image, compare.tag2)
		if err != nil {
			return err
		}

		writeDiffTableForImages(labels1, labels2, t, compare.image, compare.tag1, compare.tag2)

		if enableMarkdown {
			t.RenderMarkdown()
		} else {
			t.Render()
		}

		// write the table
		c.Log.Println(sb.String())

		// write the compare link
		rev1, rev2 := inspect.Revision(labels1), inspect.Revision(labels2)
		if rev1 != "" && rev2 != "" {
			c.Log.Println(fmt.Sprintf("%s/compare/%s..%s", inspect.BaseURL(labels1), inspect.Revision(labels1), inspect.Revision(labels2)))
		}
	}

	return nil
}

func names(in []string) []string {
	names := []string{}
	for _, n := range in {
		r, _ := ParseRepo(n)
		if !util.Contains(names, r) {
			names = append(names, r)
		}
	}
	return names
}

func locateVersion(in []string, name string) string {
	for _, n := range in {
		r, t := ParseRepo(n)
		if name == r {
			return t
		}
	}
	return ""
}
