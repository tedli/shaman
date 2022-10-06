package common

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	ErrFileOrFolderNotSpecified = errors.New("file or folder not specified")
)

func GetFile(filePath, folder string, index int) (file string, err error) {
	if filePath != "" {
		file = filePath
	} else if folder != "" {
		var entries []os.DirEntry
		if entries, err = os.ReadDir(folder); err != nil {
			return
		}
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			if index < 0 {
				file = filepath.Join(folder, "manifest.yaml")
				break
			}
			if name := entry.Name(); strings.HasSuffix(name, fmt.Sprintf("%05d.txt", index)) {
				file = filepath.Join(folder, name)
				break
			}
		}
	} else {
		err = ErrFileOrFolderNotSpecified
		return
	}
	return
}
