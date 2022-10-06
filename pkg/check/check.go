package check

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/tedli/shaman/pkg/common"
	"github.com/tedli/shaman/pkg/util"
	"io"
)

func check(folder string, out io.Writer) error {

	return common.Traversal(
		folder, func(manifest *common.Manifest, chunk *common.Chunk, chunkFile io.Reader) (err error) {
			var expect []byte
			if expect, err = hex.DecodeString(chunk.Hash); err != nil {
				return
			}
			var got []byte
			if got, err = util.Hash(chunkFile, manifest.Hash.Algorithm); err != nil {
				return
			}
			if bytes.Compare(expect, got) != 0 {
				err = common.ErrFileHashMismatch
				_, _ = fmt.Fprintf(out, "File: %s, Expect: %s, Got: %X\n",
					chunk.Name, chunk.Hash, got)
				return
			}
			return
		})
}
