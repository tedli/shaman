package generate

import "github.com/spf13/cobra"

var (
	chunkFilePath   string
	inputFileFolder string
	chunkFileIndex  int

	generateCmd = &cobra.Command{
		Use:                "generate",
		Short:              "Generate QR code pic file.",
		FParseErrWhitelist: cobra.FParseErrWhitelist{UnknownFlags: true},
		RunE: func(_ *cobra.Command, _ []string) error {
			return generate(chunkFilePath, inputFileFolder, chunkFileIndex)
		},
	}
)

func init() {
	generateCmd.PersistentFlags().StringVarP(&chunkFilePath,
		"file", "f", "", "The chunk file.")
	generateCmd.PersistentFlags().StringVarP(&inputFileFolder,
		"folder", "p", "", "The chunk file folder.")
	generateCmd.PersistentFlags().IntVarP(&chunkFileIndex,
		"index", "i", -1, "The chunk file index, use with folder. Default is manifest.yaml.")
}

func Command() *cobra.Command {
	return generateCmd
}
