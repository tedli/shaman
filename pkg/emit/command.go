package emit

import (
	"github.com/spf13/cobra"
)

var (
	chunkFilePath   string
	inputFileFolder string
	chunkFileIndex  int
	after           int

	emitCmd = &cobra.Command{
		Use:                "emit",
		Short:              "Emit keyboard typing events.",
		FParseErrWhitelist: cobra.FParseErrWhitelist{UnknownFlags: true},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return emit(chunkFilePath, inputFileFolder, chunkFileIndex, after, cmd.OutOrStderr())
		},
	}
)

func init() {
	emitCmd.PersistentFlags().StringVarP(&chunkFilePath,
		"file", "f", "", "The chunk file.")
	emitCmd.PersistentFlags().StringVarP(&inputFileFolder,
		"folder", "p", "", "The chunk file folder.")
	emitCmd.PersistentFlags().IntVarP(&chunkFileIndex,
		"index", "i", -1, "The chunk file index, use with folder. Default is manifest.yaml.")
	emitCmd.PersistentFlags().IntVarP(&after,
		"after", "d", 10, "The delay duration in second, default is 10.")
}

func Command() *cobra.Command {
	return emitCmd
}
