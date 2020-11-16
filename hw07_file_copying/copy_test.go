package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("should handle ErrOffsetExceedsFileSize error", func(t *testing.T) {
		err := Copy("testdata/input.txt", "out.txt", 1024*10, 0)
		require.Equal(t, err, ErrOffsetExceedsFileSize)
	})

	t.Run("should handle ErrUnsupportedFile error", func(t *testing.T) {
		err := Copy("some-file.txt", "out.txt", 0, 0)
		require.Equal(t, err, ErrUnsupportedFile)
	})

	t.Run("should create output file with selected limit", func(t *testing.T) {
		outFile := "out.txt"
		limit := int64(1000)
		_ = Copy("testdata/input.txt", outFile, 0, limit)
		defer func() {
			err := os.Remove(outFile)
			if err != nil {
				t.Error(err)
				return
			}
		}()

		outFileStat, err := os.Stat(outFile)

		if err != nil {
			t.Error(err)
			return
		}

		require.Equal(t, outFileStat.Size(), limit)
	})
}
