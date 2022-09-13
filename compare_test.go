package fileutils

import (
	"errors"
	"testing"

	"github.com/pthomison/errcheck"
)

func TestFileComparison(t *testing.T) {

	baseFolder := "./resources/fileComparison/"
	alpha := baseFolder + "file-alpha.txt"
	beta := baseFolder + "file-beta.txt"
	wrong := baseFolder + "file-wrong.txt"

	ExecTestFilepaths(t, alpha, beta, wrong)
}

func TestFolderComparison(t *testing.T) {

	baseFolder := "./resources/folderComparison/"
	alpha := baseFolder + "folderAlpha"
	beta := baseFolder + "folderBeta"
	wrong := baseFolder + "folderWrong"

	ExecTestFilepaths(t, alpha, beta, wrong)
}

func ExecTestFilepaths(t *testing.T, alpha string, beta string, wrong string) {

	correctOutput := CompareFilepaths(alpha, beta)
	wrongOutput := CompareFilepaths(alpha, wrong)

	if !correctOutput {
		errcheck.CheckTest(errors.New("incorrectly didn't match output"), t)
	}

	if wrongOutput {
		errcheck.CheckTest(errors.New("incorrectly matched output"), t)
	}
}
