package read

import (
	"github.com/tuotoo/qrcode"
	"os"
	"path/filepath"
	"strings"
)

func read(file string) (err error) {
	var input *os.File
	if input, err = os.Open(file); err != nil {
		return
	}
	defer func() { _ = input.Close() }()
	var matrix *qrcode.Matrix
	if matrix, err = qrcode.Decode(input); err != nil {
		return
	}
	var extension string
	if strings.HasSuffix(file, "manifest.png") {
		extension = ".yaml"
	} else {
		extension = ".txt"
	}
	file = strings.TrimSuffix(file, filepath.Ext(file)) + extension
	err = os.WriteFile(file, []byte(matrix.Content), 0644)
	return
}
