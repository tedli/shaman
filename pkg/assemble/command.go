package assemble

import "github.com/spf13/cobra"

var (
	folder string
	output string

	assembleCmd = &cobra.Command{
		Use:                "assemble",
		Short:              "Assemble chunk files to original file.",
		FParseErrWhitelist: cobra.FParseErrWhitelist{UnknownFlags: true},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return assemble(folder, output, cmd.OutOrStderr())
		},
	}
)

func init() {
	assembleCmd.PersistentFlags().StringVarP(&folder,
		"folder", "f", "", "The input folder.")
	assembleCmd.PersistentFlags().StringVarP(&output,
		"output", "o", "", "The output file.")
}

func Command() *cobra.Command {
	return assembleCmd
}
