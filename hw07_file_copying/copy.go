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
	ErrCreateFile            = errors.New("failed to create output file")
	ErrWriteFile             = errors.New("failed to write file")
	ErrSeekFile              = errors.New("failed to seek file")
	buffer                   = 1024
)

func Copy(fromPath string, toPath string, offset, limit int64) error {
	src, err := getSource(fromPath, offset, limit)
	if err != nil {
		return err
	}

	defer func() {
		err := src.file.Close()
		if err != nil {
			return
		}
	}()

	dest, err := os.Create(toPath)
	if err != nil {
		return ErrCreateFile
	}

	defer func() {
		err := dest.Close()
		if err != nil {
			return
		}
	}()

	bar := pb.Start64(src.limit)
	bufferSize := int64(math.Min(float64(buffer), float64(src.limit)))
	buf := make([]byte, bufferSize)

	for i := offset; i < offset+src.limit; i += bufferSize {
		n, err := src.file.Read(buf)
		if err != nil && !errors.Is(err, io.EOF) {
			return err
		}
		if n == 0 {
			break
		}

		bytes, err := dest.Write(buf[:n])

		if err != nil {
			return ErrWriteFile
		}

		bar.Add(bytes)
		time.Sleep(time.Millisecond * 200)
	}

	return nil
}

type source struct {
	file  *os.File
	limit int64
}

func getSource(path string, offset, limit int64) (*source, error) {
	srcStat, err := os.Stat(path)
	if err != nil || !srcStat.Mode().IsRegular() {
		return nil, ErrUnsupportedFile
	}

	srcSize := srcStat.Size()
	if srcSize < offset {
		return nil, ErrOffsetExceedsFileSize
	}
	if limit == 0 || limit+offset > srcSize {
		limit = srcSize - offset
	}

	src, err := os.Open(path)
	if err != nil {
		return nil, ErrUnsupportedFile
	}

	_, err = src.Seek(offset, 0)
	if err != nil {
		return nil, ErrSeekFile
	}

	return &source{file: src, limit: limit}, nil
}
