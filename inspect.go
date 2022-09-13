package fileutils

import (
	"fmt"
	"os"

	"github.com/pthomison/errcheck"
)

func LS(directoryPath string) {
	dir, err := os.Open(directoryPath)
	errcheck.Check(err)

	defer dir.Close()

	subNodes, err := dir.ReadDir(0)
	errcheck.Check(err)

	for _, subNode := range subNodes {
		fmt.Printf("%+v\n", subNode)
	}
}
