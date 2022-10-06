package prepare

import (
	"encoding/ascii85"
	"fmt"
	"github.com/tedli/shaman/pkg/common"
	"github.com/tedli/shaman/pkg/util"
	"github.com/ulikunitz/xz"
	"go.uber.org/multierr"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func prepare(inputFilePath, outputFolderPath, hashAlgorithm string, chunkSize int, out io.Writer) (err error) {
	var input *os.File
	if input, err = os.Open(inputFilePath); err != nil {
		return
	}
	defer func() { _ = input.Close() }()
	var hash []byte
	if hash, err = util.Hash(input, hashAlgorithm); err != nil {
		return
	}
	manifest := &common.Manifest{
		Hash: struct {
			Algorithm string
			Value     string
		}{
			Algorithm: hashAlgorithm,
			Value:     fmt.Sprintf("%X", hash),
		},
	}
	if _, err = input.Seek(0, 0); err != nil {
		return
	}
	if _, err = util.EnsureFolderExist(outputFolderPath); err != nil {
		return
	}
	filename := filepath.Base(inputFilePath)
	manifest.Name = filename
	var chunks []string
	if chunks, err = readAndEncode(
		input, outputFolderPath, filename, chunkSize, out); err != nil {
		return
	}
	chunkHashes := make([]*common.Chunk, 0, len(chunks))
	for _, c := range chunks {
		var chunkFile *os.File
		if chunkFile, err = os.Open(c); err != nil {
			return
		}
		if hash, err = util.Hash(chunkFile, hashAlgorithm); err != nil {
			return
		}
		chunkHashes = append(chunkHashes, &common.Chunk{
			Name: filepath.Base(c),
			Hash: fmt.Sprintf("%X", hash),
		})
		if err = chunkFile.Close(); err != nil {
			return
		}
	}
	manifest.Chunks = chunkHashes
	var manifestFile *os.File
	if manifestFile, err = os.Create(filepath.Join(outputFolderPath, "manifest.yaml")); err != nil {
		return
	}
	defer func() { _ = manifestFile.Close() }()
	err = yaml.NewEncoder(manifestFile).Encode(manifest)
	return
}

func readAndEncode(input io.Reader, output, filename string, chunkSize int, out io.Writer) (chunks []string, err error) {
	filename = strings.TrimSuffix(filename, filepath.Ext(filename))
	chunkWriter := util.NewChunkWriter(chunkSize, func(index int) (io.WriteCloser, error) {
		if chunks == nil {
			chunks = make([]string, 0)
		}
		chunkFilename := fmt.Sprintf("%s_%05d.txt", filename, index)
		path := filepath.Join(output, chunkFilename)
		chunks = append(chunks, path)
		_, _ = fmt.Fprintf(out, "Generate chunk file: %s\n", chunkFilename)
		return os.Create(path)
	})
	text := ascii85.NewEncoder(chunkWriter)
	var compress io.WriteCloser
	if compress, err = xz.NewWriter(text); err != nil {
		return
	}
	errs := make([]error, 0, 3)
	if _, err = io.Copy(compress, input); err != nil {
		errs = append(errs, err)
	}
	if err = compress.Close(); err != nil {
		errs = append(errs, err)
	}
	if err = text.Close(); err != nil {
		errs = append(errs, err)
	}
	if err = chunkWriter.Close(); err != nil {
		errs = append(errs, err)
	}
	if el := len(errs); el > 0 {
		if el == 1 {
			err = errs[0]
		} else {
			err = multierr.Combine(errs...)
		}
	}
	return
}
