package main

import (
	"errors"
	"io"
	"math"
	"os"
	"time"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	buffer                   = 1024
)

func Copy(fromPath string, toPath string, offset, limit int64) error {
	sourceStat, err := os.Stat(fromPath)
	if err != nil || !sourceStat.Mode().IsRegular() {
		return ErrUnsupportedFile
	}

	sourceSize := sourceStat.Size()
	if sourceSize < offset {
		return ErrOffsetExceedsFileSize
	}
	if limit == 0 || limit+offset > sourceSize {
		limit = sourceSize - offset
	}

	source, err := os.Open(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}

	defer func() {
		err := source.Close()
		if err != nil {
			return
		}
	}()

	_, err = source.Seek(offset, 0)
	if err != nil {
		return err
	}

	dest, err := os.Create(toPath)
	if err != nil {
		return err
	}

	defer func() {
		err := dest.Close()
		if err != nil {
			return
		}
	}()

	bar := pb.Start64(limit)

	bufferSize := int64(math.Min(float64(buffer), float64(limit)))
	buf := make([]byte, bufferSize)

	for i := offset; i < offset+limit; i += bufferSize {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		bytes, err := dest.Write(buf[:n])

		if err != nil {
			return err
		}

		bar.Add(bytes)
		time.Sleep(time.Millisecond * 200)
	}

	return nil
}
