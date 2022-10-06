package util

import (
	"errors"
	"os"
)

var (
	ErrPathExistNotFolder = errors.New("path exist not folder")
)

func EnsureFolderExist(path string) (create bool, err error) {
	var info os.FileInfo
	if info, err = os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			err = nil
			create = true
		} else {
			return
		}
	}
	if !create && !info.IsDir() {
		err = ErrPathExistNotFolder
		return
	}
	if !create {
		return
	}
	err = os.MkdirAll(path, 0644)
	return
}
