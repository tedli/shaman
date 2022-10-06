package generate

import (
	"github.com/skip2/go-qrcode"
	"github.com/tedli/shaman/pkg/common"
	"os"
	"path/filepath"
	"strings"
)

func generate(filePath, folder string, index int) (err error) {
	var file string
	if file, err = common.GetFile(filePath, folder, index); err != nil {
		return
	}
	var content []byte
	if content, err = os.ReadFile(file); err != nil {
		return
	}
	filename := filepath.Base(file)
	filename = strings.TrimSuffix(file, filepath.Ext(filename)) + ".png"
	err = qrcode.WriteFile(string(content), qrcode.Medium, 512, filename)
	return
}
