package fileutils

import (
	"bytes"
	"os"
	"path/filepath"

	"github.com/pthomison/errcheck"
)

func CompareFilepaths(folderAlpha string, folderBeta string) bool {
	// DRYing funcs
	isDir := func(node *os.File) bool {
		stat, err := node.Stat()
		errcheck.Check(err)
		is := stat.IsDir()
		return is
	}

	index := func(node *os.File) map[string]os.DirEntry {
		files, err := node.ReadDir(0)
		errcheck.Check(err)

		index := make(map[string]os.DirEntry)

		for _, v := range files {
			index[v.Name()] = v
		}

		return index
	}

	alphaNode := OpenFilepath(folderAlpha)
	defer alphaNode.Close()
	betaNode := OpenFilepath(folderBeta)
	defer betaNode.Close()

	// If either node is nil, it means one of the folders doesn't have a file; return false
	// if alphaNode == nil || betaNode == nil {
	// 	return false
	// }

	alphaIsDir := isDir(alphaNode)
	betaIsDir := isDir(betaNode)

	// If one node is a dir && the other isn't, return false
	if (alphaIsDir && !betaIsDir) || (!alphaIsDir && betaIsDir) {
		return false
	}

	// If both nodes are files, load their contents and compare
	if !alphaIsDir && !betaIsDir {
		alphaBytes := ReadFile(alphaNode)
		betaBytes := ReadFile(betaNode)

		return bytes.Equal(alphaBytes, betaBytes)
	}

	// If both nodes are directorys, list their contents && recursivly call this function against their contents
	// return all the resusive answers logically AND'd
	if alphaIsDir && betaIsDir {
		alphaIndex := index(alphaNode)
		betaIndex := index(betaNode)

		result := true

		for name, _ := range alphaIndex {

			if betaIndex[name] == nil {
				return false
			}

			alphaSubpath := filepath.Join(folderAlpha, name)
			betaSubpath := filepath.Join(folderBeta, name)

			result = result && CompareFilepaths(alphaSubpath, betaSubpath)
		}

		return result

	}

	return false
}
