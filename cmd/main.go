package main

import (
	"flag"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tedli/shaman/pkg/assemble"
	"github.com/tedli/shaman/pkg/check"
	"github.com/tedli/shaman/pkg/emit"
	"github.com/tedli/shaman/pkg/generate"
	"github.com/tedli/shaman/pkg/prepare"
	"github.com/tedli/shaman/pkg/read"
	"github.com/tedli/shaman/pkg/util"
	"github.com/tedli/shaman/pkg/version"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:                "shaman",
		Short:              "A data transfer tool.",
		Long:               "A data transfer tool, emitting keyboard typing events, and generate, read QR code.",
		SilenceUsage:       true,
		FParseErrWhitelist: cobra.FParseErrWhitelist{UnknownFlags: true},
	}
)

func init() {
	rootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	rootCmd.AddCommand(
		prepare.Command(),
		emit.Command(),
		assemble.Command(),
		generate.Command(),
		read.Command(),
		check.Command(),
		util.Command(),
		version.Command())
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(rootCmd.OutOrStderr(), "Error: %s\n", err)
		os.Exit(-1)
	}
}
