package assemble

import (
	"bytes"
	"encoding/ascii85"
	"encoding/hex"
	"fmt"
	"github.com/tedli/shaman/pkg/common"
	"github.com/tedli/shaman/pkg/util"
	"github.com/ulikunitz/xz"
	"io"
	"os"
)

func assemble(folder, output string, out io.Writer) error {
	file, e := os.Create(output)
	if e != nil {
		return e
	}
	defer func() { _ = file.Close() }()
	var m *common.Manifest
	if e = common.Traversal(
		folder, func(manifest *common.Manifest, _ *common.Chunk, chunkFile io.Reader) (err error) {
			if m == nil {
				m = manifest
			}
			var input io.Reader
			if input, err = xz.NewReader(ascii85.NewDecoder(chunkFile)); err != nil {
				return
			}
			if _, err = io.Copy(file, input); err != nil {
				return
			}
			return
		}); e != nil {
		return e
	}
	if e = file.Sync(); e != nil {
		return e
	}
	if _, e = file.Seek(0, 0); e != nil {
		return e
	}
	var got []byte
	if got, e = util.Hash(file, m.Hash.Algorithm); e != nil {
		return e
	}
	var expect []byte
	if expect, e = hex.DecodeString(m.Hash.Value); e != nil {
		return e
	}
	if bytes.Compare(expect, got) != 0 {
		_, _ = fmt.Fprintf(out, "File: %s, Expect: %s, Got: %X\n",
			m.Name, m.Hash.Value, got)
		return common.ErrFileHashMismatch
	}
	return nil
}
