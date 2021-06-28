package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

var (
	blockCopy                int64 = 512
	ErrUnsupportedFile             = errors.New("unsupported file")
	ErrOffsetExceedsFileSize       = errors.New("offset exceeds file size")
	ErrOpeningFile                 = errors.New("error opening file")
	ErrCreatingFile                = errors.New("error creating file")
	ErrCopyingFile                 = errors.New("error copying to file")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	start := time.Now()

	var file *os.File
	file, err := os.OpenFile(fromPath, os.O_RDONLY, 0644)
	if err != nil {
		return ErrOpeningFile
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return ErrUnsupportedFile
	}
	fileSize := fileInfo.Size()
	if offset > fileSize {
		return ErrOffsetExceedsFileSize
	}
	if (limit == 0) || (limit+offset) > fileSize {
		limit = fileSize - offset
	}
	seek := offset

	fileNew, err := os.Create(toPath)
	if err != nil {
		return ErrCreatingFile
	}
	defer fileNew.Close()

	bar := NewBar(limit)

	for {
		file.Seek(seek, 0)
		if (limit + offset) < (seek + blockCopy) {
			blockCopy = limit + offset - seek
		}
		c, err := io.CopyN(fileNew, file, blockCopy)

		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return ErrCopyingFile
		}

		seek += c

		barData := seek - offset
		bar.Update(barData)

		if (limit + offset) <= seek {
			break
		}
	}

	bar.Finish()

	fmt.Printf("copied %d bytes, %v \n", limit, time.Since(start))

	return nil
}
