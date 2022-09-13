package fileutils

import (
	"github.com/pthomison/errcheck"
	"gopkg.in/yaml.v3"
)

type UnstructureYamlData map[interface{}]interface{}

func ReadYamlFilepath(path string) UnstructureYamlData {
	yamlBytes := ReadFilepath(path)

	yamlData := make(UnstructureYamlData)

	err := yaml.Unmarshal(yamlBytes, &yamlData)
	errcheck.Check(err)

	return yamlData
}
