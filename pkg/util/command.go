package util

import (
	"bytes"
	"encoding/ascii85"
	"github.com/spf13/cobra"
	"github.com/ulikunitz/xz"
	"io"
	"strings"
)

var (
	utilCmd = &cobra.Command{
		Use:                "util",
		Short:              "Utilities for decompress or decode.",
		FParseErrWhitelist: cobra.FParseErrWhitelist{UnknownFlags: true},
	}
	xzdCmd = &cobra.Command{
		Use:                "xz",
		Short:              "Decompress xz compressed content.",
		FParseErrWhitelist: cobra.FParseErrWhitelist{UnknownFlags: true},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return xzDecompress(cmd.InOrStdin(), cmd.OutOrStdout())
		},
	}
	b85Cmd = &cobra.Command{
		Use:                "ascii",
		Short:              "Decode ascii85 encoded content.",
		FParseErrWhitelist: cobra.FParseErrWhitelist{UnknownFlags: true},
		RunE: func(cmd *cobra.Command, args []string) error {
			return base85Decode(strings.Join(args, ""), cmd.InOrStdin(), cmd.OutOrStdout())
		},
	}
)

func init() {
	utilCmd.AddCommand(xzdCmd, b85Cmd)
}

func Command() *cobra.Command {
	return utilCmd
}

func xzDecompress(input io.Reader, output io.Writer) (err error) {
	var xzReader io.Reader
	if xzReader, err = xz.NewReader(input); err != nil {
		return
	}
	_, err = io.Copy(output, xzReader)
	return
}

func base85Decode(content string, input io.Reader, output io.Writer) (err error) {
	var in io.Reader
	if content != "" {
		in = bytes.NewBufferString(content)
	} else {
		in = input
	}
	decoder := ascii85.NewDecoder(in)
	_, err = io.Copy(output, decoder)
	return
}
