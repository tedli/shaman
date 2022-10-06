package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	Release  = "v0.0.1"
	CommitID = "dev"
	BuiltAt  = "dev"

	readCmd = &cobra.Command{
		Use:   "version",
		Short: "Prints out build version information",
		RunE: func(cmd *cobra.Command, _ []string) error {
			_, err := fmt.Fprintf(cmd.OutOrStdout(),
				"Release: %s, CommitID: %s, BuiltAt: %s\n", Release, CommitID, BuiltAt)
			return err
		},
	}
)

func Command() *cobra.Command {
	return readCmd
}
