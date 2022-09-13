package fileutils

import (
	"errors"
	"os"

	"github.com/pthomison/errcheck"
)

func ReadFilepath(path string) []byte {
	data, err := os.ReadFile(path)
	errcheck.Check(err)

	return data
}

func ReadFile(file *os.File) []byte {
	var fileBytes []byte

	fileStat, err := file.Stat()
	errcheck.Check(err)
	fileSize := fileStat.Size()

	fileBytes = make([]byte, fileSize)

	readSize, err := file.Read(fileBytes)
	errcheck.Check(err)

	if readSize != int(fileSize) {
		errcheck.Check(errors.New("Read bytes does not equal file size"))
	}

	return fileBytes
}

func OpenFilepath(path string) *os.File {
	file, err := os.Open(path)
	errcheck.Check(err)

	return file
}
