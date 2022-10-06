package prepare

import (
	"github.com/spf13/cobra"
)

var (
	inputFilePath    string
	outputFolderPath string
	hashAlgorithm    string
	chunkSize        int

	prepareCmd = &cobra.Command{
		Use:                "prepare",
		Short:              "Compress and encode file to several chunks for generating keyboard events next.",
		FParseErrWhitelist: cobra.FParseErrWhitelist{UnknownFlags: true},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return prepare(inputFilePath, outputFolderPath, hashAlgorithm, chunkSize, cmd.OutOrStdout())
		},
	}
)

func init() {
	prepareCmd.PersistentFlags().StringVarP(&inputFilePath,
		"input", "i", "", "The input file.")
	prepareCmd.PersistentFlags().StringVarP(&outputFolderPath,
		"output", "o", "", "The output folder.")
	prepareCmd.PersistentFlags().StringVarP(&hashAlgorithm,
		"hash", "a", "SHA1", "The hash algorithm, default is SHA1, available SHA1, MD5.")
	prepareCmd.PersistentFlags().IntVarP(&chunkSize,
		"size", "s", 1024, "The chunk file size in bytes, default is 1024 (1K).")
}

func Command() *cobra.Command {
	return prepareCmd
}
