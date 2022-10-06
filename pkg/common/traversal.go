package common

import (
	"go.uber.org/multierr"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"path/filepath"
)

func Traversal(folder string, action func(*Manifest, *Chunk, io.Reader) error) (err error) {
	var manifestFile io.ReadCloser
	if manifestFile, err = os.Open(filepath.Join(folder, "manifest.yaml")); err != nil {
		return
	}
	defer func() { _ = manifestFile.Close() }()
	var manifest Manifest
	if err = yaml.NewDecoder(manifestFile).Decode(&manifest); err != nil {
		return
	}
	for _, chunk := range manifest.Chunks {
		var chunkFile io.ReadCloser
		if chunkFile, err = os.Open(filepath.Join(folder, chunk.Name)); err != nil {
			return
		}
		if err = action(&manifest, chunk, chunkFile); err != nil {
			if e := chunkFile.Close(); e != nil {
				err = multierr.Combine(err, e)
			}
			return
		}
		if err = chunkFile.Close(); err != nil {
			return
		}
	}
	return
}
