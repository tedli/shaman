package check

import "github.com/spf13/cobra"

var (
	folder string

	checkCmd = &cobra.Command{
		Use:                "check",
		Short:              "Check the content of transferred files by emitting keyboard events.",
		FParseErrWhitelist: cobra.FParseErrWhitelist{UnknownFlags: true},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return check(folder, cmd.OutOrStderr())
		},
	}
)

func init() {
	checkCmd.PersistentFlags().StringVarP(&folder,
		"folder", "f", "", "The input folder.")
}

func Command() *cobra.Command {
	return checkCmd
}
