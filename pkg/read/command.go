package read

import "github.com/spf13/cobra"

var (
	input string

	readCmd = &cobra.Command{
		Use:                "read",
		Short:              "Read QR code and write content to chunk file.",
		FParseErrWhitelist: cobra.FParseErrWhitelist{UnknownFlags: true},
		RunE: func(_ *cobra.Command, _ []string) error {
			return read(input)
		},
	}
)

func init() {
	readCmd.PersistentFlags().StringVarP(&input,
		"input", "i", "", "The input file.")
}

func Command() *cobra.Command {
	return readCmd
}
