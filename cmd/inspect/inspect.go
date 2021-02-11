package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"

	"github.com/garethjevans/inspect/pkg/util"

	"github.com/garethjevans/inspect/pkg/cmd"

	"github.com/garethjevans/inspect/pkg/version"

	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra/doc"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// Version is dynamically set by the toolchain or overridden by the Makefile.
var Version = version.Version

var Verbose bool

// BuildDate is dynamically set at build time in the Makefile.
var BuildDate = version.BuildDate

var versionOutput = ""

func init() {
	if strings.Contains(Version, "dev") {
		if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "(devel)" {
			Version = info.Main.Version
		}
	}
	Version = strings.TrimPrefix(Version, "v")
	if BuildDate == "" {
		RootCmd.Version = Version
	} else {
		RootCmd.Version = fmt.Sprintf("%s (%s)", Version, BuildDate)
	}
	versionOutput = fmt.Sprintf("inspect version %s\n", RootCmd.Version)
	RootCmd.AddCommand(versionCmd)
	RootCmd.SetVersionTemplate(versionOutput)

	RootCmd.AddCommand(docsCmd)

	RootCmd.PersistentFlags().Bool("help", false, "Show help for command")
	RootCmd.PersistentFlags().BoolVarP(&Verbose, "debug", "v", false, "Debug Output")

	RootCmd.Flags().Bool("version", false, "Show version")

	RootCmd.SetFlagErrorFunc(func(cmd *cobra.Command, err error) error {
		if err == pflag.ErrHelp {
			return err
		}
		return &FlagError{Err: err}
	})

	RootCmd.AddCommand(cmd.NewImageCmd())
	RootCmd.AddCommand(cmd.NewBuildArgsCmd())
	RootCmd.AddCommand(cmd.NewLabelsCmd())
	RootCmd.AddCommand(cmd.NewDiffCmd())

	RootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		if Verbose {
			logrus.SetLevel(logrus.DebugLevel)
		}
	}

	c := completionCmd
	c.Flags().StringP("shell", "s", "bash", "Shell type: {bash|zsh|fish|powershell}")
	RootCmd.AddCommand(c)
}

// FlagError is the kind of error raised in flag processing.
type FlagError struct {
	Err error
}

// Error.
func (fe FlagError) Error() string {
	return fe.Err.Error()
}

// Unwrap FlagError.
func (fe FlagError) Unwrap() error {
	return fe.Err
}

// RootCmd is the entry point of command-line execution.
var RootCmd = &cobra.Command{
	Use:   "inspect",
	Short: "inspect a docker image",
	Long:  `a simple CLI to inspect a docker image and try and determine the origin of the container`,

	SilenceErrors: false,
	SilenceUsage:  false,
}

var versionCmd = &cobra.Command{
	Use:    "version",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(versionOutput)
	},
}

var docsCmd = &cobra.Command{
	Use:    "docs",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		RootCmd.DisableAutoGenTag = true

		err := doc.GenMarkdownTree(RootCmd, "./docs")
		if err != nil {
			panic(err)
		}
	},
}

var completionCmd = &cobra.Command{
	Use:    "completion",
	Hidden: true,
	Short:  "Generate shell completion scripts",
	Long: `Generate shell completion scripts for GitHub CLI commands.

The output of this command will be computer code and is meant to be saved to a
file or immediately evaluated by an interactive shell.

For example, for bash you could add this to your '~/.bash_profile':

	eval "$(gh completion -s bash)"

When installing GitHub CLI through a package manager, however, it's possible that
no additional shell configuration is necessary to gain completion support. For
Homebrew, see <https://docs.brew.sh/Shell-Completion>
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		shellType, err := cmd.Flags().GetString("shell")
		if err != nil {
			return err
		}

		if shellType == "" {
			shellType = "bash"
		}

		switch shellType {
		case "bash":
			return RootCmd.GenBashCompletion(cmd.OutOrStdout())
		case "zsh":
			return RootCmd.GenZshCompletion(cmd.OutOrStdout())
		case "fish":
			return RootCmd.GenFishCompletion(cmd.OutOrStdout(), true)
		case "powershell":
			return RootCmd.GenPowerShellCompletion(cmd.OutOrStdout())
		default:
			return fmt.Errorf("unsupported shell type %q", shellType)
		}
	},
}

func main() {
	logrus.SetFormatter(util.NewTextFormat())
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
