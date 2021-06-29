package main

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("error ErrOpeningFile", func(t *testing.T) {
		fromPath := "testdata/input"
		toPath := "out.txt"

		var offset int64 = 0
		var limit int64 = 1000
		err := Copy(fromPath, toPath, offset, limit)

		require.Truef(t, errors.Is(err, ErrOpeningFile), "actual err - %v", err)
	})
	t.Run("error ErrOffsetExceedsFileSize", func(t *testing.T) {
		fromPath := "testdata/input.txt"
		toPath := "out.txt"
		fileInfo, err := os.Stat(fromPath)
		if err != nil {
			t.Error("no open file fromPatch")
		}

		offset := fileInfo.Size() + 1
		var limit int64 = 1
		err = Copy(fromPath, toPath, offset, limit)

		require.Truef(t, errors.Is(err, ErrOffsetExceedsFileSize), "actual err - %v", err)
	})
}
