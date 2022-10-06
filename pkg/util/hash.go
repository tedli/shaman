package util

import (
	"crypto/md5"
	"crypto/sha1"
	"errors"
	"hash"
	"io"
)

type ctor func() hash.Hash

var (
	algorithmCtorMapping = map[string]ctor{
		"SHA1": sha1.New,
		"MD5":  md5.New,
	}

	ErrHashAlgorithmNotSupported = errors.New("hash algorithm not supported")
)

func Hash(input io.Reader, algorithm string) (result []byte, err error) {
	create, supported := algorithmCtorMapping[algorithm]
	if !supported {
		err = ErrHashAlgorithmNotSupported
		return
	}
	hasher := create()
	if _, err = io.Copy(hasher, input); err != nil {
		return
	}
	result = hasher.Sum(nil)
	return
}
