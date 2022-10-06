package util

import (
	"bytes"
	"io"
)

func NewChunkWriter(chunkSize int, ctor func(int) (io.WriteCloser, error)) io.WriteCloser {
	return &chunkWriter{
		chunkSize:       chunkSize,
		chunkWriterCtor: ctor,
		buffer:          new(bytes.Buffer),
		index:           0,
	}
}

type chunkWriter struct {
	chunkSize       int
	chunkWriterCtor func(int) (io.WriteCloser, error)
	buffer          *bytes.Buffer
	currentWriter   io.WriteCloser
	index           int
}

func (cw *chunkWriter) Write(data []byte) (n int, err error) {
	if n, err = cw.buffer.Write(data); err != nil {
		return
	}
	for cw.buffer.Len() >= cw.chunkSize {
		if cw.currentWriter == nil {
			if cw.currentWriter, err = cw.chunkWriterCtor(cw.index); err != nil {
				return
			}
		}
		buffer := make([]byte, cw.chunkSize)
		if _, err = cw.buffer.Read(buffer); err != nil {
			return
		}
		if _, err = io.Copy(cw.currentWriter, bytes.NewReader(buffer)); err != nil {
			return
		}
		if err = cw.currentWriter.Close(); err != nil {
			return
		}
		cw.index++
		cw.currentWriter = nil
		if rest := cw.buffer.Len(); rest > 0 && rest < cw.chunkSize {
			if rest, err = cw.buffer.Read(buffer); err != nil {
				return
			}
			cw.buffer.Reset()
			cw.buffer.Write(buffer[:rest])
		} else if rest == 0 {
			cw.buffer.Reset()
		}
	}
	return
}

func (cw *chunkWriter) Close() (err error) {
	var rest int
	if rest = cw.buffer.Len(); rest == 0 {
		return
	}
	if cw.currentWriter == nil {
		if cw.currentWriter, err = cw.chunkWriterCtor(cw.index); err != nil {
			return
		}
	}
	buffer := make([]byte, rest)
	if rest, err = cw.buffer.Read(buffer); err != nil {
		return
	}
	if _, err = io.Copy(cw.currentWriter, bytes.NewReader(buffer)); err != nil {
		return
	}
	if err = cw.currentWriter.Close(); err != nil {
		return
	}
	cw.index = 0
	cw.currentWriter = nil
	return
}
