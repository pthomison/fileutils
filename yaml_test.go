package fileutils

import (
	"errors"
	"testing"

	"github.com/pthomison/errcheck"
)

func TestYamlLoad(t *testing.T) {

	correctData := make(UnstructureYamlData)
	correctData["one"] = 1
	correctData["two"] = "two"
	correctData["three"] = true

	yamlFilepath := "./resources/yamlLoad/test.yaml"
	yamlData := ReadYamlFilepath(yamlFilepath)

	if (yamlData["one"] != 1) ||
		(yamlData["two"] != "two") ||
		(yamlData["three"] != true) {

		errcheck.CheckTest(errors.New("Loaded YAML data didn't match test data"), t)
	}
}
