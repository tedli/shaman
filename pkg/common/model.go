package common

import "errors"

type Chunk struct {
	Name string
	Hash string
}

type Manifest struct {
	Name string
	Hash struct {
		Algorithm string
		Value     string
	}
	Chunks []*Chunk
}

var (
	ErrFileHashMismatch = errors.New("file hash mismatch")
)
